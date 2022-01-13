package models

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math"
	"math/big"
	"openseasync/common/utils"
	"openseasync/database"
	"openseasync/logs"
	"strconv"
	"time"
)

const ZeroAddress = "0x0000000000000000000000000000000000000000"

// InsertOpenSeaAsset query Aseets through opensea API and insert
func InsertOpenSeaAsset(assets *OwnerAsset, user string) error {
	db := database.GetMongoClient()
	refreshTime := int(time.Now().Unix())

	for _, v := range assets.Assets {

		var (
			asset = Asset{
				UserMetamaskID:    user,
				Title:             v.Name,
				ImageURL:          v.ImageURL,
				ImagePreviewURL:   v.ImagePreviewURL,
				ImageThumbnailURL: v.ImageThumbnailURL,
				Description:       v.Description,
				ContractAddress:   v.AssetContract.Address,
				TokenId:           v.TokenID,
				NumSales:          v.NumSales,
				Owner:             v.Owner.Address,
				OwnerName:         v.Owner.User.Username,
				OwnerImgURL:       v.Owner.ProfileImgURL,
				Creator:           v.Creator.Address,
				CreatorName:       v.Creator.User.Username,
				CreatorImgURL:     v.Creator.ProfileImgURL,
				CollectionID:      v.Collection.Slug,
				TokenMetadata:     v.TokenMetadata,
				RefreshTime:       refreshTime,
			}
			contract = Contract{
				Address:      v.AssetContract.Address,
				ContractName: v.AssetContract.Name,
				ContractType: v.AssetContract.AssetContractType,
				Symbol:       v.AssetContract.Symbol,
				SchemaName:   v.AssetContract.SchemaName,
				TotalSupply:  v.AssetContract.TotalSupply,
				Description:  v.AssetContract.Description,
			}
			traits             []Trait
			assetTopOwnerships []AssetsTopOwnership
		)

		for _, v1 := range v.Traits {
			var trait = Trait{
				UserMetamaskID:  user,
				ContractAddress: v.AssetContract.Address,
				TokenId:         v.TokenID,
				TraitType:       v1.TraitType,
				Value:           v1.Value,
				DisplayType:     v1.DisplayType,
				MaxValue:        v1.MaxValue,
				TraitCount:      v1.TraitCount,
				OrderBy:         v1.Order,
				RefreshTime:     refreshTime,
			}
			traits = append(traits, trait)
		}
		asset.Traits = traits
		if v.SellOrders != nil {
			asset.SellOrders.CreateDate = v.SellOrders[0].CreatedDate
			asset.SellOrders.ClosingDate = v.SellOrders[0].ClosingDate
			asset.SellOrders.CurrentPrice = v.SellOrders[0].CurrentPrice
			asset.SellOrders.PayTokenContract.Symbol = v.SellOrders[0].PaymentTokenContract.Symbol
			asset.SellOrders.PayTokenContract.ImageURL = v.SellOrders[0].PaymentTokenContract.ImageURL
			asset.SellOrders.PayTokenContract.EthPrice = v.SellOrders[0].PaymentTokenContract.EthPrice
			asset.SellOrders.PayTokenContract.UsdPrice = v.SellOrders[0].PaymentTokenContract.UsdPrice
		}

		time.Sleep(time.Second * 2)
		// If the number of requests is too many, a 429 error code will be thrown
		resp, err := utils.RequestOpenSeaSingleAsset(v.AssetContract.Address, v.TokenID)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}

		// insert top_owner_ships
		var autoAsset AutoAsset
		if err = json.Unmarshal(resp, &autoAsset); err != nil {
			logs.GetLogger().Error(err)
			return err
		}

		for _, v2 := range autoAsset.TopOwnerships {
			var assetsTopOwnership = AssetsTopOwnership{
				UserMetamaskID:  user,
				ContractAddress: v.AssetContract.Address,
				TokenId:         v.TokenID,
				Owner:           v2.Owner.Address,
				ProfileImgURL:   v2.Owner.ProfileImgURL,
				Quantity:        v2.Quantity,
				RefreshTime:     refreshTime,
			}
			assetTopOwnerships = append(assetTopOwnerships, assetsTopOwnership)
		}
		asset.AssetsTopOwnerships = assetTopOwnerships

		// insert assets
		count, err := db.Collection("assets").CountDocuments(
			context.TODO(),
			bson.M{"userMetamaskId": user, "contract_address": v.AssetContract.Address, "token_id": v.TokenID,
				"is_delete": 0})
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
		if count == 0 {
			if _, err = db.Collection("assets").InsertOne(context.TODO(), &asset); err != nil {
				logs.GetLogger().Error(err)
				return err
			}
		} else {
			var sellOrders bson.M
			if v.SellOrders != nil {
				sellOrders = bson.M{
					"create_date":   v.SellOrders[0].CreatedDate,
					"closing_date":  v.SellOrders[0].ClosingDate,
					"current_price": v.SellOrders[0].CurrentPrice,
					"pay_token_contract": bson.M{
						"symbol":    v.SellOrders[0].PaymentTokenContract.Symbol,
						"image_url": v.SellOrders[0].PaymentTokenContract.ImageURL,
						"eth_price": v.SellOrders[0].PaymentTokenContract.EthPrice,
						"usd_price": v.SellOrders[0].PaymentTokenContract.UsdPrice,
					},
				}
			}
			// update
			if _, err = db.Collection("assets").UpdateOne(
				context.TODO(),
				bson.M{"userMetamaskId": user, "contract_address": v.AssetContract.Address, "token_id": v.TokenID, "is_delete": 0},
				bson.M{"$set": bson.M{"title": v.Name, "image_url": v.ImageURL, "image_preview_url": v.ImagePreviewURL,
					"image_thumbnail_url": v.ImageThumbnailURL, "description": v.Description, "refresh_time": refreshTime,
					"traits": traits, "assets_top_ownerships": assetTopOwnerships, "sell_orders": sellOrders}}); err != nil {
				logs.GetLogger().Error(err)
				return err
			}
		}

		// insert user
		var userModel = User{
			Id:             0,
			UserMetamaskID: user,
			Username:       v.Owner.User.Username,
			UserImgURL:     v.Owner.ProfileImgURL,
		}
		if v.Owner.Address == ZeroAddress && user == v.Creator.Address {
			userModel.Username = v.Creator.User.Username
			userModel.UserImgURL = v.Creator.ProfileImgURL
		}
		if err := insertUsers(db, user, userModel); err != nil {
			logs.GetLogger().Error(err)
			return err
		}

		// insert order
		if err := insertOrders(db, v.AssetContract.Address, v.TokenID, autoAsset); err != nil {
			logs.GetLogger().Error(err)
			return err
		}

		// insert contract
		if err := insertContract(db, v.AssetContract.Address, &contract); err != nil {
			logs.GetLogger().Error(err)
			return err
		}

		// insert transaction
		time.Sleep(time.Second * 2)
		if err := insertTransaction(db, user, v.AssetContract.Address, v.TokenID); err != nil {
			logs.GetLogger().Error(err)
			return err
		}

	}

	// Delete opensea deleted asset
	if err := deleteAsset(db, user, refreshTime); err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	return nil
}

// FindAssetByOwner find assets by owner
func FindAssetByOwner(user string, collectionId interface{}, page, pageSize int64) (map[string]interface{}, error) {
	var (
		assets     []*Asset
		assetsList []map[string]interface{}
		result     = make(map[string]interface{})
	)
	db := database.GetMongoClient()

	condition := bson.M{"userMetamaskId": user, "is_delete": 0}
	if collectionId != nil {
		condition["collectionId"] = collectionId
	}

	total, err := db.Collection("assets").CountDocuments(context.TODO(), condition)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	totalPage := total / pageSize
	if total%pageSize != 0 {
		totalPage++
	}
	opts := options.Find().SetSkip((page - 1) * pageSize).SetLimit(pageSize)
	cursor, err := db.Collection("assets").Find(context.TODO(), condition, opts)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	if err = cursor.All(context.TODO(), &assets); err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	for _, v := range assets {
		var (
			collection    Collection
			contract      Contract
			itemActivitys []ItemActivity
			orders        []Orders
			data          map[string]interface{}
		)
		err = db.Collection("collections").FindOne(
			context.TODO(), bson.M{"userMetamaskId": user, "id": v.CollectionID}).Decode(&collection)
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}

		err = db.Collection("contracts").FindOne(
			context.TODO(), bson.M{"address": v.ContractAddress}).Decode(&contract)
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}

		cursor, err := db.Collection("item_activitys").Find(
			context.TODO(), bson.M{"userMetamaskId": user, "contract_address": v.ContractAddress, "token_id": v.TokenId, "is_delete": 0})
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}
		if err = cursor.All(context.TODO(), &itemActivitys); err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}

		cursor, err = db.Collection("orders").Find(
			context.TODO(), bson.M{"contract_address": v.ContractAddress, "token_id": v.TokenId, "is_delete": 0})
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}
		if err = cursor.All(context.TODO(), &orders); err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}

		bytes, err := json.Marshal(v)
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}
		if err = json.Unmarshal(bytes, &data); err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}

		data["contract"] = contract
		data["collection"] = collection
		data["item_activitys"] = itemActivitys
		data["orders"] = orders
		assetsList = append(assetsList, data)
	}

	result["data"] = assetsList
	result["metadata"] = map[string]int64{"page": page, "pageSize": pageSize, "total": total, "totalPage": totalPage}

	return result, nil
}

// DeleteAssetByTokenID delete asset by tokenId
func DeleteAssetByTokenID(user, contractAddress, tokenID string) error {
	db := database.GetMongoClient()
	// delete assets
	if _, err := db.Collection("assets").UpdateMany(
		context.TODO(),
		bson.M{"userMetamaskId": user, "contract_address": contractAddress, "token_id": tokenID},
		bson.M{"$set": bson.M{"is_delete": 1}}); err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	// delete item_activitys
	if _, err := db.Collection("item_activitys").UpdateMany(
		context.TODO(),
		bson.M{"userMetamaskId": user, "contract_address": contractAddress, "token_id": tokenID},
		bson.M{"$set": bson.M{"is_delete": 1}}); err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	// delete orders
	if _, err := db.Collection("orders").UpdateMany(
		context.TODO(),
		bson.M{"contract_address": contractAddress, "token_id": tokenID},
		bson.M{"$set": bson.M{"is_delete": 1}}); err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	return nil
}

// insert contract
func insertContract(db *mongo.Database, contractAddress string, contract *Contract) error {
	count, err := db.Collection("contracts").
		CountDocuments(context.TODO(), bson.M{"address": contractAddress})
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	if count == 0 {
		if _, err = db.Collection("contracts").InsertOne(context.TODO(), contract); err != nil {
			logs.GetLogger().Error(err)
			return err
		}
	}
	return nil
}

// delete asset
func deleteAsset(db *mongo.Database, user string, refreshTime int) error {
	if _, err := db.Collection("assets").UpdateMany(
		context.TODO(),
		bson.M{"userMetamaskId": user, "refresh_time": bson.M{"$lt": refreshTime}, "is_delete": 0},
		bson.M{"$set": bson.M{"is_delete": 1}}); err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	return nil
}

// insert transaction
func insertTransaction(db *mongo.Database, user, contractAddress, tokenId string) error {
	// If the number of requests is too many, a 429 error code will be thrown
	resp, err := utils.RequestOpenSeaEvent(contractAddress, tokenId)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	var event Event
	if err = json.Unmarshal(resp, &event); err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	// insert item_activitys
	for _, v := range event.AssetEvents {
		var itemActivity = ItemActivity{
			Id:               v.ID,
			CollectibleId:    v.Asset.ID,
			CollectibleName:  v.Asset.Name,
			CollectionId:     v.CollectionSlug,
			CollectionName:   v.Asset.Collection.Name,
			ContractAddress:  contractAddress,
			TokenId:          tokenId,
			BidAmount:        v.BidAmount,
			CreateDate:       v.CreatedDate,
			Price:            v.TotalPrice,
			SellerMetamaskId: v.Seller.Address,
			SellerName:       v.Seller.User.Username,
			SellerImgURL:     v.Seller.ProfileImgURL,
			BuyerMetamaskId:  v.WinnerAccount.Address,
			BuyerName:        v.WinnerAccount.User.Username,
			BuyerImgURL:      v.WinnerAccount.ProfileImgURL,
			TradeType:        v.EventType,
			Quantity:         v.Quantity,
			Transaction:      v.Transaction,
		}

		if v.PaymentToken.UsdPrice != nil {
			itemActivity.PayTokenContract = PayTokenContract{
				Symbol:   v.PaymentToken.Symbol,
				ImageURL: v.PaymentToken.ImageURL,
				EthPrice: v.PaymentToken.EthPrice,
				UsdPrice: v.PaymentToken.UsdPrice.(string),
			}
			usdPrice, _ := strconv.ParseFloat(v.PaymentToken.UsdPrice.(string), 64)
			price, _ := strconv.ParseFloat(v.TotalPrice, 64)
			n := new(big.Int)
			n.Mul(big.NewInt(int64(usdPrice)), big.NewInt(int64(price)))
			n.Div(n, big.NewInt(int64(math.Pow10(18))))
			itemActivity.PriceInUsd = n.String()
		}

		count, err := db.Collection("item_activitys").CountDocuments(
			context.TODO(),
			bson.M{"id": v.ID, "contract_address": contractAddress, "token_id": tokenId, "is_delete": 0})
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
		if count == 0 {
			if _, err = db.Collection("item_activitys").InsertOne(context.TODO(), &itemActivity); err != nil {
				logs.GetLogger().Error(err)
				return err
			}
		} else {
			// update
			itemActivityByte, err := bson.Marshal(itemActivity)
			if err != nil {
				logs.GetLogger().Error(err)
				return err
			}
			var tmpItemActivity bson.M
			if err := bson.Unmarshal(itemActivityByte, &tmpItemActivity); err != nil {
				logs.GetLogger().Error(err)
				return err
			}
			if _, err = db.Collection("item_activitys").UpdateOne(
				context.TODO(),
				bson.M{"contract_address": contractAddress, "token_id": tokenId, "is_delete": 0},
				bson.M{"$set": tmpItemActivity}); err != nil {
				logs.GetLogger().Error(err)
				return err
			}
		}
	}
	return nil
}

// insert orders
func insertOrders(db *mongo.Database, contractAddress, tokenId string, autoAsset AutoAsset) error {
	for _, v := range autoAsset.Orders {
		var orders = Orders{
			ContractAddress: contractAddress,
			TokenId:         tokenId,
			CreateDate:      v.CreatedDate,
			ClosingDate:     v.ClosingDate,
			ExpirationTime:  v.ExpirationTime,
			ListingTime:     v.ListingTime,
			OrderHash:       v.OrderHash,
			CurrentPrice:    v.CurrentPrice,
			CurrentBounty:   v.CurrentBounty,
			BasePrice:       v.BasePrice,
			PaymentToken:    v.PaymentToken,
			Target:          v.Target,
		}
		orders.Metadata.ID = v.Metadata.Asset.ID
		orders.Metadata.Address = v.Metadata.Asset.Address
		orders.Metadata.Quantity = v.Metadata.Asset.Quantity
		orders.Metadata.Schema = v.Metadata.Schema
		orders.Maker.UserName = v.Maker.User.Username
		orders.Maker.ProfileImgURL = v.Maker.ProfileImgURL
		orders.Maker.Address = v.Maker.Address
		orders.Taker.UserName = v.Taker.User.Username
		orders.Taker.Address = v.Taker.Address
		orders.Taker.ProfileImgURL = v.Taker.ProfileImgURL
		orders.PayTokenContract.Symbol = v.PaymentTokenContract.Symbol
		orders.PayTokenContract.ImageURL = v.PaymentTokenContract.ImageURL
		orders.PayTokenContract.EthPrice = v.PaymentTokenContract.EthPrice
		orders.PayTokenContract.UsdPrice = v.PaymentTokenContract.UsdPrice
		count, err := db.Collection("orders").
			CountDocuments(context.TODO(), bson.M{"order_hash": v.OrderHash})
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
		if count == 0 {
			if _, err = db.Collection("orders").InsertOne(context.TODO(), &orders); err != nil {
				logs.GetLogger().Error(err)
				return err
			}
		} else {
			ordersByte, err := bson.Marshal(orders)
			if err != nil {
				logs.GetLogger().Error(err)
				return err
			}
			var tmpOrders bson.M
			if err := bson.Unmarshal(ordersByte, &tmpOrders); err != nil {
				logs.GetLogger().Error(err)
				return err
			}

			if _, err = db.Collection("assets").UpdateOne(
				context.TODO(),
				bson.M{"order_hash": v.OrderHash}, bson.M{"$set": tmpOrders}); err != nil {
				logs.GetLogger().Error(err)
				return err
			}
		}

	}
	return nil
}

// insert users
func insertUsers(db *mongo.Database, userAddress string, user User) error {
	count, err := db.Collection("users").
		CountDocuments(context.TODO(), bson.M{"userMetamaskId": userAddress})
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	if count == 0 {
		if _, err = db.Collection("users").InsertOne(context.TODO(), user); err != nil {
			logs.GetLogger().Error(err)
			return err
		}
	} else {
		// update
		if _, err = db.Collection("users").UpdateOne(
			context.TODO(),
			bson.M{"userMetamaskId": userAddress},
			bson.M{"$set": bson.M{"userName": user.Username, "userImgURL": user.UserImgURL}}); err != nil {
			logs.GetLogger().Error(err)
			return err
		}
	}
	return nil
}
