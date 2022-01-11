package models

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"openseasync/common/utils"
	"openseasync/database"
	"openseasync/logs"
	"time"
)

// InsertOpenSeaAsset query Aseets through opensea API and insert
func InsertOpenSeaAsset(assets *OwnerAsset, user string) error {
	db := database.GetMongoClient()
	refreshTime := int(time.Now().Unix())

	for _, v := range assets.Assets {
		owner := user
		if v.Owner.Address != "0x0000000000000000000000000000000000000000" {
			owner = v.Owner.Address
		}

		var (
			asset = Asset{
				UserAddress:       user,
				Title:             v.Name,
				ImageURL:          v.ImageURL,
				ImagePreviewURL:   v.ImagePreviewURL,
				ImageThumbnailURL: v.ImageThumbnailURL,
				Description:       v.Description,
				ContractAddress:   v.AssetContract.Address,
				TokenId:           v.TokenID,
				NumSales:          v.NumSales,
				Owner:             owner,
				OwnerImgURL:       v.Owner.ProfileImgURL,
				Creator:           v.Creator.Address,
				CreatorImgURL:     v.Creator.ProfileImgURL,
				Slug:              v.Collection.Slug,
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
				UserAddress:     user,
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
		var autoAsset AutoAsset
		if err = json.Unmarshal(resp, &autoAsset); err != nil {
			logs.GetLogger().Error(err)
			return err
		}

		for _, v2 := range autoAsset.TopOwnerships {
			var assetsTopOwnership = AssetsTopOwnership{
				UserAddress:     user,
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

		count, err := db.Collection("assets").CountDocuments(
			context.TODO(),
			bson.M{"user_address": user, "contract_address": v.AssetContract.Address, "token_id": v.TokenID,
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
				bson.M{"user_address": user, "contract_address": v.AssetContract.Address, "token_id": v.TokenID, "is_delete": 0},
				bson.M{"$set": bson.M{"title": v.Name, "image_url": v.ImageURL, "image_preview_url": v.ImagePreviewURL,
					"image_thumbnail_url": v.ImageThumbnailURL, "description": v.Description, "refresh_time": refreshTime,
					"traits": traits, "assets_top_ownerships": assetTopOwnerships, "sell_orders": sellOrders}}); err != nil {
				logs.GetLogger().Error(err)
				return err
			}
		}

		// insert contract
		count, err = db.Collection("contracts").
			CountDocuments(context.TODO(), bson.M{"address": v.AssetContract.Address})
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
		if count == 0 {
			if _, err = db.Collection("contracts").InsertOne(context.TODO(), &contract); err != nil {
				logs.GetLogger().Error(err)
				return err
			}
		}

		// Insert transaction
		time.Sleep(time.Second * 2)
		// If the number of requests is too many, a 429 error code will be thrown
		resp, err = utils.RequestOpenSeaEvent(v.AssetContract.Address, v.TokenID)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
		var event Event
		if err = json.Unmarshal(resp, &event); err != nil {
			logs.GetLogger().Error(err)
			return err
		}

		for _, v3 := range event.AssetEvents {
			var itemActivity = ItemActivity{
				Id:                  v3.ID,
				UserAddress:         user,
				ContractAddress:     v.AssetContract.Address,
				TokenId:             v.TokenID,
				BidAmount:           v3.BidAmount,
				CreateDate:          v3.CreatedDate,
				TotalPrice:          v3.TotalPrice,
				Seller:              v3.Seller.Address,
				SellerProfileImgURL: v3.Seller.ProfileImgURL,
				Winner:              v3.WinnerAccount.Address,
				WinnerProfileImgURL: v3.WinnerAccount.ProfileImgURL,
				EventType:           v3.EventType,
				Transaction:         v3.Transaction,
			}
			count, err := db.Collection("item_activitys").CountDocuments(
				context.TODO(),
				bson.M{"user_address": user, "contract_address": v.AssetContract.Address, "token_id": v.TokenID, "id": v3.ID,
					"is_delete": 0})
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
				if _, err = db.Collection("item_activitys").UpdateOne(
					context.TODO(),
					bson.M{"user_address": user, "contract_address": v.AssetContract.Address, "token_id": v.TokenID, "is_delete": 0},
					bson.M{"$set": bson.M{
						"bid_amount": v3.BidAmount, "create_date": v3.CreatedDate, "total_price": v3.TotalPrice,
						"seller": v3.Seller.Address, "seller_profile_img_url": v3.Seller.ProfileImgURL, "event_type": v3.EventType,
						"winner": v3.WinnerAccount.Address, "winner_profile_img_url": v3.WinnerAccount.ProfileImgURL,
						"transaction": v3.Transaction}}); err != nil {
					logs.GetLogger().Error(err)
					return err
				}
			}
		}
	}

	// Delete opensea deleted asset
	if _, err := db.Collection("assets").UpdateMany(
		context.TODO(),
		bson.M{"user_address": user, "refresh_time": bson.M{"$lt": refreshTime}, "is_delete": 0},
		bson.M{"$set": bson.M{"is_delete": 1}}); err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	return nil
}

// FindAssetByOwner find assets by owner
func FindAssetByOwner(user string) ([]map[string]interface{}, error) {
	var (
		assets     []*Asset
		assetsList []map[string]interface{}
	)
	db := database.GetMongoClient()
	cursor, err := db.Collection("assets").Find(context.TODO(), bson.M{"user_address": user, "is_delete": 0})
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
			result        map[string]interface{}
		)
		err = db.Collection("collections").FindOne(
			context.TODO(), bson.M{"user_address": user, "slug": v.Slug}).Decode(&collection)
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
			context.TODO(), bson.M{"user_address": user, "contract_address": v.ContractAddress, "token_id": v.TokenId, "is_delete": 0})
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}
		if err = cursor.All(context.TODO(), &itemActivitys); err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}

		bytes, err := json.Marshal(v)
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}
		if err = json.Unmarshal(bytes, &result); err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}

		result["contract"] = contract
		result["collection"] = collection
		result["item_activitys"] = itemActivitys
		assetsList = append(assetsList, result)
	}

	return assetsList, nil
}

// FindWorksBySlug find assets by collection
func FindWorksBySlug(user, slug string) ([]*Asset, error) {
	var assets []*Asset
	db := database.GetMongoClient()

	cursor, err := db.Collection("assets").Find(
		context.TODO(), bson.M{"user_address": user, "slug": slug, "is_delete": 0})
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	if err = cursor.All(context.TODO(), &assets); err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return assets, nil
}

// DeleteAssetByTokenID delete asset by tokenId
func DeleteAssetByTokenID(user, contractAddress, tokenID string) error {
	db := database.GetMongoClient()
	if _, err := db.Collection("assets").UpdateMany(
		context.TODO(),
		bson.M{"user_address": user, "contract_address": contractAddress, "token_id": tokenID},
		bson.M{"$set": bson.M{"is_delete": 1}}); err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	if _, err := db.Collection("item_activitys").UpdateMany(
		context.TODO(),
		bson.M{"user_address": user, "contract_address": contractAddress, "token_id": tokenID},
		bson.M{"$set": bson.M{"is_delete": 1}}); err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	return nil
}
