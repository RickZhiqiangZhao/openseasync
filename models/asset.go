package models

import (
	"context"
	"encoding/json"
	uuid2 "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	refreshTime := int64(time.Now().Unix())

	for _, v := range assets.Assets {
		uuidOrder, _ := uuid2.NewUUID()
		var (
			asset = Asset{
				Id:                 v.ID,
				UserMetamaskID:     user,
				Name:               v.Name,
				CoverImageUrl:      v.ImageURL,
				CoverPreviewUrl:    v.ImagePreviewURL,
				ThumbnailUrl:       v.ImageThumbnailURL,
				FileUrl:            v.TokenMetadata,
				Description:        v.Description,
				ContractAddress:    v.AssetContract.Address,
				CollectibleTokenId: v.TokenID,
				NumSales:           v.NumSales,
				OwnerMetamaskId:    v.Owner.Address,
				OwnerName:          v.Owner.User.Username,
				OwnerImgURL:        v.Owner.ProfileImgURL,
				CreatorMetamaskId:  v.Creator.Address,
				CreatorName:        v.Creator.User.Username,
				CreatorImgUrl:      v.Creator.ProfileImgURL,
				CollectionID:       v.Collection.Slug,
				TokenMetadata:      v.TokenMetadata,
				RefreshTime:        refreshTime,
				Status:             "onHold",
				NumOfCopies:        1,
				TotalCopies:        1,
				RecordId:           uuidOrder.String(),
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
			sellOrders := v.SellOrders[0]
			asset.Price = sellOrders.CurrentPrice
			asset.StartTime = utils.ParseTime(sellOrders.CreatedDate)
			if v.SellOrders[0].ClosingDate != "" {
				asset.EndTime = utils.ParseTime(sellOrders.ClosingDate)
			}
			asset.Status = "onSale"
			if sellOrders.ClosingExtendable == true {
				asset.Status = "onAuction"
			}

			asset.SellOrders.CreateDate = sellOrders.CreatedDate
			asset.SellOrders.ClosingDate = sellOrders.ClosingDate
			asset.SellOrders.CurrentPrice = sellOrders.CurrentPrice
			asset.SellOrders.PayTokenContract.Symbol = sellOrders.PaymentTokenContract.Symbol
			asset.SellOrders.PayTokenContract.ImageURL = sellOrders.PaymentTokenContract.ImageURL
			asset.SellOrders.PayTokenContract.EthPrice = sellOrders.PaymentTokenContract.EthPrice
			asset.SellOrders.PayTokenContract.UsdPrice = sellOrders.PaymentTokenContract.UsdPrice
		}
		// insert transaction
		time.Sleep(time.Second * 2)
		createDate, err := insertTransaction(db, v.AssetContract.Address, v.TokenID)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
		asset.CreateDate = createDate

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
				ProfileImgUrl:   v2.Owner.ProfileImgURL,
				Quantity:        v2.Quantity,
				RefreshTime:     refreshTime,
			}
			assetTopOwnerships = append(assetTopOwnerships, assetsTopOwnership)
		}
		asset.AssetsTopOwnerships = assetTopOwnerships

		// insert assets
		count, err := db.Collection("assets").CountDocuments(
			context.TODO(),
			bson.M{"userMetamaskId": user, "contractAddress": v.AssetContract.Address, "collectibleTokenId": v.TokenID,
				"isDelete": 0})
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
			// update
			assetByte, err := bson.Marshal(asset)
			if err != nil {
				logs.GetLogger().Error(err)
				return err
			}
			var tmpAssetByte bson.M
			if err := bson.Unmarshal(assetByte, &tmpAssetByte); err != nil {
				logs.GetLogger().Error(err)
				return err
			}
			if _, err = db.Collection("assets").UpdateOne(
				context.TODO(),
				bson.M{"userMetamaskId": user, "contractAddress": v.AssetContract.Address, "collectibleTokenId": v.TokenID, "isDelete": 0},
				bson.M{"$set": tmpAssetByte}); err != nil {
				logs.GetLogger().Error(err)
				return err
			}
		}

		// insert user
		uuid, _ := uuid2.NewUUID()

		var userModel = User{
			Id:               uuid.String(),
			UserMetamaskID:   user,
			Username:         v.Owner.User.Username,
			AvatarUrl:        v.Owner.ProfileImgURL,
			PersonalPageLink: v.ExternalLink,
			Discord:          v.Collection.DiscordURL,
			Telegram:         v.Collection.TelegramURL,
			Twitter:          v.Collection.TwitterUsername,
			Instagram:        v.Collection.InstagramUsername,
		}
		if v.Owner.Address == ZeroAddress && user == v.Creator.Address {
			userModel.Username = v.Creator.User.Username
			userModel.AvatarUrl = v.Creator.ProfileImgURL
		}
		if err := insertUsers(db, user, userModel); err != nil {
			logs.GetLogger().Error(err)
			return err
		}

		// insert order
		if err := insertOrders(db, v.ID, autoAsset, uuidOrder.String()); err != nil {
			logs.GetLogger().Error(err)
			return err
		}

		// insert contract
		if err := insertContract(db, v.AssetContract.Address, &contract); err != nil {
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

// FindAssetSearchByOwner find assets by owner
func FindAssetSearchByOwner(collectionId string, param Params) (map[string]interface{}, error) {
	var (
		assets = make([]bson.M, 0)
		result = make(map[string]interface{})
	)
	db := database.GetMongoClient()
	var (
		sort   bson.M
		status string
	)
	switch param.SortBy {
	case 0:
		sort = bson.M{"createDate": -1}
	case 1:
		sort = bson.M{"createDate": 1}
	case 2:
		sort = bson.M{"price": 1}
	case 3:
		sort = bson.M{"price": -1}
	case 4:
		sort = bson.M{"viewCounts": -1}
	case 5:
		sort = bson.M{"viewCounts": 1}
	case 6:
		sort = bson.M{"endTime": -1}
	}
	if param.Status == 0 {
		status = "onSale"
	} else {
		status = "onAuction"
	}

	cond := mongo.Pipeline{
		{{"$match", bson.M{"collectionId": collectionId, "status": status, "name": primitive.Regex{Pattern: param.Field, Options: "i"}, "isDelete": 0}}},
		{{
			"$addFields", bson.M{"price": bson.M{"$cond": bson.M{
				"if":   bson.M{"$ne": bson.A{"$price", ""}},
				"then": bson.M{"$convert": bson.M{"input": "$price", "to": "double"}},
				"else": 0,
			},
			}},
		}},
		{{"$match", bson.M{"price": bson.M{"$gte": param.MinPrice, "$lte": param.MaxPrice}}}},
	}
	cursor, err := db.Collection("assets").Aggregate(context.TODO(), cond)
	if err != nil && err != mongo.ErrNoDocuments {
		logs.GetLogger().Error(err)
		return nil, err
	}
	total := int64(cursor.RemainingBatchLength())
	totalPage := total / param.PageSize
	if total%param.PageSize != 0 {
		totalPage++
	}
	pipe := mongo.Pipeline{
		{{"$sort", sort}},
		{{"$skip", (param.Page - 1) * param.PageSize}},
		{{"$limit", param.PageSize}},
		{{"$project",
			bson.M{
				"_id": 0, "id": 1, "coverImageUrl": 1, "name": 1, "creatorId ": 1, "creatorMetamaskId": 1,
				"creatorName": 1, "likesCount": 1, "viewsCount": 1, "numOfCopies": 1, "totalCopies": 1, "status": 1,
				"ownerUserId": 1, "ownerMetamaskId": 1, "createDate": 1, "endTime": 1,
				"price": bson.M{"$cond": bson.M{
					"if":   bson.M{"$ne": bson.A{"$price", 0}},
					"then": "$price",
					"else": nil}},
			},
		}},
	}
	cond = append(cond, pipe...)
	cursor, err = db.Collection("assets").Aggregate(context.TODO(), cond)
	if err != nil && err != mongo.ErrNoDocuments {
		logs.GetLogger().Error(err)
		return nil, err
	}
	if err = cursor.All(context.TODO(), &assets); err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	result["data"] = assets
	result["metadata"] = map[string]int64{"page": param.Page, "pageSize": param.PageSize, "total": total, "totalPage": totalPage}

	return result, nil
}

// FindAssetByGeneralInfoCollectibleId find assets by collectibleId
func FindAssetByGeneralInfoCollectibleId(collectibleId int64) (map[string]interface{}, error) {
	var (
		assets = make([]bson.M, 0)
		result = make(map[string]interface{})
	)
	db := database.GetMongoClient()
	opts := options.Find().SetProjection(
		bson.D{
			{"_id", 0},
			{"id", 1},
			{"collectibleName", 1},
			{"collectionId", 1},
			{"collectionName", 1},
			{"creatorId ", 1},
			{"creatorMetamaskId", 1},
			{"creatorName", 1},
			{"creatorPersonalSite", 1},
			{"description", 1},
			{"fileUrl", 1},
			{"ownerId", 1},
			{"ownerMetamaskId", 1},
			{"ownerName", 1},
			{"price", 1},
			{"status", 1},
			{"thumbnailUrl", 1},
			{"collectibleTokenId", 1},
			{"recordId", 1},
			{"startTime", 1},
			{"endTime", 1},
		})
	cursor, err := db.Collection("assets").Find(context.TODO(), bson.M{"id": collectibleId}, opts)
	if err != nil && err != mongo.ErrNoDocuments {
		logs.GetLogger().Error(err)
		return nil, err
	}
	if err = cursor.All(context.TODO(), &assets); err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	result["data"] = assets
	return result, nil
}

func FindAssetOfferRecordsByCollectibleId(collectibleId int64) ([]bson.M, error) {
	var orders = make([]bson.M, 0)
	db := database.GetMongoClient()
	opts := options.Find().SetProjection(
		bson.D{
			{"_id", 0},
			{"id", 1},
			{"auctionUserId", 1},
			{"auctionMetamaskId", 1},
			{"auctionUserName", 1},
			{"price", 1},
			{"bidTime", 1},
		})
	cursor, err := db.Collection("orders").Find(context.TODO(), bson.M{"collectibleId": collectibleId}, opts)
	if err != nil && err != mongo.ErrNoDocuments {
		logs.GetLogger().Error(err)
		return nil, err
	}
	if err = cursor.All(context.TODO(), &orders); err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return orders, nil
}

// FindAssetOtherByCollection find assets other
func FindAssetOtherByCollection(collectibleId int64) (map[string]interface{}, error) {
	var (
		asset  Asset
		assets = make([]bson.M, 0)
		result = make(map[string]interface{})
	)
	db := database.GetMongoClient()
	opts := options.Find().SetSort(bson.M{"createDate": -1}).SetLimit(4).SetProjection(
		bson.D{
			{"_id", 0},
			{"id", 1},
			{"coverImageUrl", 1},
			{"collectibleName", 1},
			{"creatorId", 1},
			{"creatorName", 1},
			{"creatorMetamaskId", 1},
			{"price", 1},
			{"likesCount", 1},
			{"viewsCount", 1},
			{"numOfCopies", 1},
			{"totalCopies", 1},
			{"status", 1},
			{"ownerUserId", 1},
			{"ownerMetamaskId", 1},
			{"createDate", 1},
		})
	err := db.Collection("assets").
		FindOne(context.TODO(), bson.M{"id": collectibleId}).Decode(&asset)
	if err != nil && err != mongo.ErrNoDocuments {
		logs.GetLogger().Error(err)
		return nil, err
	}
	collectionId := asset.CollectionID
	cursor, err := db.Collection("assets").Find(context.TODO(), bson.M{"collectionId": collectionId}, opts)
	if err = cursor.All(context.TODO(), &assets); err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	result["data"] = assets
	return result, nil
}

// DeleteAssetByTokenID delete asset by tokenId
func DeleteAssetByTokenID(user, contractAddress, tokenID string) error {
	db := database.GetMongoClient()
	// delete assets
	if _, err := db.Collection("assets").UpdateMany(
		context.TODO(),
		bson.M{"userMetamaskId": user, "contractAddress": contractAddress, "collectibleTokenId": tokenID},
		bson.M{"$set": bson.M{"isDelete": 1}}); err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	// delete item_activitys
	if _, err := db.Collection("item_activitys").UpdateMany(
		context.TODO(),
		bson.M{"userMetamaskId": user, "contractAddress": contractAddress, "tokenId": tokenID},
		bson.M{"$set": bson.M{"isDelete": 1}}); err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	// delete orders
	if _, err := db.Collection("orders").UpdateMany(
		context.TODO(),
		bson.M{"contract_address": contractAddress, "tokenId": tokenID},
		bson.M{"$set": bson.M{"isDelete": 1}}); err != nil {
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
func deleteAsset(db *mongo.Database, user string, refreshTime int64) error {
	if _, err := db.Collection("assets").UpdateMany(
		context.TODO(),
		bson.M{"userMetamaskId": user, "refreshTime": bson.M{"$lt": refreshTime}, "isDelete": 0},
		bson.M{"$set": bson.M{"isDelete": 1}}); err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	return nil
}

// insert transaction
func insertTransaction(db *mongo.Database, contractAddress, tokenId string) (int64, error) {
	// If the number of requests is too many, a 429 error code will be thrown
	resp, err := utils.RequestOpenSeaEvent(contractAddress, tokenId)
	if err != nil {
		logs.GetLogger().Error(err)
		return 0, err
	}
	var event Event
	if err = json.Unmarshal(resp, &event); err != nil {
		logs.GetLogger().Error(err)
		return 0, err
	}

	var createData int64
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
			CreateDate:       utils.ParseTime(v.CreatedDate),
			Price:            v.TotalPrice,
			SellerMetamaskId: v.Seller.Address,
			SellerName:       v.Seller.User.Username,
			SellerImgUrl:     v.Seller.ProfileImgURL,
			BuyerMetamaskId:  v.WinnerAccount.Address,
			BuyerName:        v.WinnerAccount.User.Username,
			BuyerImgURL:      v.WinnerAccount.ProfileImgURL,
			TradeType:        v.EventType,
			Quantity:         v.Quantity,
			Transaction:      v.Transaction,
		}
		if v.EventType == "created" {
			if createData == 0 || createData > utils.ParseTime(v.CreatedDate) {
				createData = utils.ParseTime(v.CreatedDate)
			}
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
			bson.M{"id": v.ID, "contractAddress": contractAddress, "tokenId": tokenId, "isDelete": 0})
		if err != nil {
			logs.GetLogger().Error(err)
			return 0, err
		}
		if count == 0 {
			if _, err = db.Collection("item_activitys").InsertOne(context.TODO(), &itemActivity); err != nil {
				logs.GetLogger().Error(err)
				return 0, err
			}
		} else {
			// update
			itemActivityByte, err := bson.Marshal(itemActivity)
			if err != nil {
				logs.GetLogger().Error(err)
				return 0, err
			}
			var tmpItemActivity bson.M
			if err := bson.Unmarshal(itemActivityByte, &tmpItemActivity); err != nil {
				logs.GetLogger().Error(err)
				return 0, err
			}
			if _, err = db.Collection("item_activitys").UpdateOne(
				context.TODO(),
				bson.M{"contractAddress": contractAddress, "tokenId": tokenId, "isDelete": 0},
				bson.M{"$set": tmpItemActivity}); err != nil {
				logs.GetLogger().Error(err)
				return 0, err
			}
		}
	}
	return createData, nil
}

// insert orders
func insertOrders(db *mongo.Database, collectibleId int, autoAsset AutoAsset, uuid string) error {
	for _, v := range autoAsset.Orders {
		var orders = Orders{
			UUID:              uuid,
			Id:                v.OrderHash,
			CollectibleId:     collectibleId,
			CreateDate:        v.CreatedDate,
			ClosingDate:       v.ClosingDate,
			BidTime:           v.CreatedDate,
			AuctionMetamaskId: v.Maker.Address,
			AuctionUserName:   v.Maker.User.Username,
			CurrentBounty:     v.CurrentBounty,
			Price:             v.CurrentPrice,
			BasePrice:         v.BasePrice,
		}
		orders.PayTokenContract.Symbol = v.PaymentTokenContract.Symbol
		orders.PayTokenContract.ImageURL = v.PaymentTokenContract.ImageURL
		orders.PayTokenContract.EthPrice = v.PaymentTokenContract.EthPrice
		orders.PayTokenContract.UsdPrice = v.PaymentTokenContract.UsdPrice
		count, err := db.Collection("orders").
			CountDocuments(context.TODO(), bson.M{"orderHash": v.OrderHash})
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
				bson.M{"orderHash": v.OrderHash}, bson.M{"$set": tmpOrders}); err != nil {
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
			bson.M{"$set": bson.M{"userName": user.Username, "avatarUrl": user.AvatarUrl}}); err != nil {
			logs.GetLogger().Error(err)
			return err
		}
	}
	return nil
}
