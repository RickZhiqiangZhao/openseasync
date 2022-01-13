package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"openseasync/database"
	"openseasync/logs"
)

// FindItemActivityByCollectionId find item_activity by collection_id
func FindItemActivityByCollectionId(collectionId string, page, pageSize int64) (map[string]interface{}, error) {
	var (
		itemActivitys []bson.M
		result        = make(map[string]interface{})
	)
	db := database.GetMongoClient()

	total, err := db.Collection("item_activitys").CountDocuments(context.TODO(), bson.M{"collectionId": collectionId, "is_delete": 0})
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	totalPage := total / pageSize
	if total%pageSize != 0 {
		totalPage++
	}
	opts := options.Find().SetSkip((page - 1) * pageSize).SetLimit(pageSize).SetProjection(
		bson.D{
			{"id", 1},
			{"tradeType", 1},
			{"price", 1},
			{"priceInUsd ", 1},
			{"collectibleId", 1},
			{"collectibleName", 1},
			{"quantity ", 1},
			{"buyerId", 1},
			{"buyerMetamaskId", 1},
			{"buyerName", 1},
			{"sellerId", 1},
			{"sellerMetamaskId", 1},
			{"sellerName", 1},
		})
	cursor, err := db.Collection("item_activitys").Find(context.TODO(), bson.M{"collectionId": collectionId, "is_delete": 0}, opts)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	if err = cursor.All(context.TODO(), &itemActivitys); err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	result["data"] = itemActivitys
	result["metadata"] = map[string]int64{"page": page, "pageSize": pageSize, "total": total, "totalPage": totalPage}

	return result, nil
}
