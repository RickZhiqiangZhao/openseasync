package models

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"openseasync/database"
	"openseasync/logs"
	"time"
)

var CONNOT_DELETE_COLLECTION_ERR = errors.New("Cannot delete a collection that has an asset")

// InsertOpenSeaCollection find collection through opensea API and insert
func InsertOpenSeaCollection(collections *OwnerCollection, user string) error {
	db := database.GetMongoClient()
	refreshTime := int(time.Now().Unix())
	for _, v := range collections.Collections {
		var collection = Collection{
			ID:                 v.Slug,
			UserMetamaskID:     user,
			CollectionName:     v.Name,
			Description:        v.Description,
			UserCoverUrl:       v.BannerImageURL,
			CoverImageUrl:      v.ImageURL,
			CoverLargeImageUrl: v.LargeImageURL,
			CreateDate:         v.CreatedDate,
			RefreshTime:        refreshTime,
			OwnersCount:        v.Stats.NumOwners,
			ItemsCount:         int(v.Stats.TotalSupply),
			TotalVolume:        v.Stats.TotalVolume,
			FloorPrice:         v.Stats.FloorPrice,
		}
		count, err := db.Collection("collections").
			CountDocuments(context.TODO(), bson.M{"userMetamaskId": user, "id": v.Slug, "isDelete": 0})
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
		if count == 0 {
			if _, err = db.Collection("collections").InsertOne(context.TODO(), &collection); err != nil {
				logs.GetLogger().Error(err)
				return err
			}
		} else {
			// update
			collectionByte, err := bson.Marshal(collection)
			if err != nil {
				logs.GetLogger().Error(err)
				return err
			}
			var tmpCollection bson.M
			if err := bson.Unmarshal(collectionByte, &tmpCollection); err != nil {
				logs.GetLogger().Error(err)
				return err
			}
			if _, err = db.Collection("collections").UpdateOne(
				context.TODO(),
				bson.M{"userMetamaskId": user, "id": v.Slug, "isDelete": 0},
				bson.M{"$set": tmpCollection}); err != nil {
				logs.GetLogger().Error(err)
				return err
			}
		}
	}

	if _, err := db.Collection("collections").UpdateMany(
		context.TODO(),
		bson.M{"userMetamaskId": user, "refreshTime": bson.M{"$lt": refreshTime}, "isDelete": 0},
		bson.M{"$set": bson.M{"isDelete": 1}}); err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	return nil

}

// FindCollectionByUserMetamaskID find collections by usermetamaskid
func FindCollectionByUserMetamaskID(usermetamaskid string, page, pageSize int64) (map[string]interface{}, error) {
	var (
		collections []bson.M
		result      = make(map[string]interface{})
	)
	db := database.GetMongoClient()
	total, err := db.Collection("collections").CountDocuments(context.TODO(), bson.M{"userMetamaskId": usermetamaskid, "isDelete": 0})
	if err != nil && err != mongo.ErrNoDocuments {
		logs.GetLogger().Error(err)
		return nil, err
	}
	totalPage := total / pageSize
	if total%pageSize != 0 {
		totalPage++
	}

	pipe := mongo.Pipeline{
		{{"$match", bson.M{"userMetamaskId": usermetamaskid, "isDelete": 0}}},
		{{"$skip", (page - 1) * pageSize}},
		{{"$limit", pageSize}},
		{{"$lookup", bson.M{
			"from":     "users",
			"let":      bson.M{"userMetamaskId": "$userMetamaskId"},
			"pipeline": bson.A{bson.M{"$match": bson.M{"$expr": bson.M{"$eq": bson.A{"$userMetamaskId", "$$userMetamaskId"}}}}},
			"as":       "user_item",
		}}},
		{{
			"$addFields", bson.M{"user_item": bson.M{"$arrayElemAt": bson.A{"$user_item", 0}}},
		}},
		{{
			"$addFields", bson.M{"userId": "$user_item.id", "userName": "$user_item.userName", "avatarUrl": "$user_item.avatarUrl"},
		}},
		{{"$project",
			bson.M{
				"_id": 0, "id": 1, "userId": 1, "userMetamaskId": 1, "coverImageUrl ": 1, "avatarUrl": 1, "userName": 1,
				"collectionName": 1, "description": 1,
			},
		}},
	}
	cursor, err := db.Collection("collections").Aggregate(context.TODO(), pipe)
	if err != nil && err != mongo.ErrNoDocuments {
		logs.GetLogger().Error(err)
		return nil, err
	}
	if err = cursor.All(context.TODO(), &collections); err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	result["data"] = collections
	result["metadata"] = map[string]int64{"page": page, "pageSize": pageSize, "total": total, "totalPage": totalPage}

	return result, nil
}

// FindCollectionByCollectionID find collections by collectionId
func FindCollectionByCollectionID(collectionId string, page, pageSize int64) (map[string]interface{}, error) {
	var (
		collections []bson.M
		result      = make(map[string]interface{})
	)
	db := database.GetMongoClient()
	total, err := db.Collection("collections").CountDocuments(context.TODO(), bson.M{"id": collectionId, "isDelete": 0})
	if err != nil && err != mongo.ErrNoDocuments {
		logs.GetLogger().Error(err)
		return nil, err
	}
	totalPage := total / pageSize
	if total%pageSize != 0 {
		totalPage++
	}
	pipe := mongo.Pipeline{
		{{"$match", bson.M{"id": collectionId, "isDelete": 0}}},
		{{"$skip", (page - 1) * pageSize}},
		{{"$limit", pageSize}},
		{{"$lookup", bson.M{
			"from":     "users",
			"let":      bson.M{"userMetamaskId": "$userMetamaskId"},
			"pipeline": bson.A{bson.M{"$match": bson.M{"$expr": bson.M{"$eq": bson.A{"$userMetamaskId", "$$userMetamaskId"}}}}},
			"as":       "user_item",
		}}},
		{{
			"$addFields", bson.M{"user_item": bson.M{"$arrayElemAt": bson.A{"$user_item", 0}}},
		}},
		{{
			"$addFields", bson.M{"userId": "$user_item.id", "userName": "$user_item.userName", "avatarUrl": "$user_item.avatarUrl"},
		}},
		{{"$project",
			bson.M{"_id": 0, "id": 1, "userId": 1, "userMetamaskId": 1, "userCoverUrl": 1, "avatarUrl": 1,
				"userName": 1, "itemsCount": 1, "ownersCount": 1, "floorPrice": 1, "highestPrice": 1,
				"collectionName": 1, "likesCount": 1, "viewsCount": 1, "description": 1}}},
	}
	cursor, err := db.Collection("collections").Aggregate(context.TODO(), pipe)
	if err != nil && err != mongo.ErrNoDocuments {
		logs.GetLogger().Error(err)
		return nil, err
	}
	if err = cursor.All(context.TODO(), &collections); err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	result["data"] = collections
	result["metadata"] = map[string]int64{"page": page, "pageSize": pageSize, "total": total, "totalPage": totalPage}

	return result, nil
}

// FindUserMediaByUserId find user media by userId
func FindUserMediaByUserId(userId string) (bson.M, error) {
	var user bson.M
	db := database.GetMongoClient()
	if err := db.Collection("users").FindOne(context.TODO(), bson.M{"id": userId}).Decode(&user); err != nil && err != mongo.ErrNoDocuments {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return user, nil
}

// DeleteCollectionByCollectionId delete empty collection
func DeleteCollectionByCollectionId(user, slug string) error {
	db := database.GetMongoClient()
	row, err := db.Collection("assets").CountDocuments(
		context.TODO(), bson.M{"userMetamaskId": user, "collectionId": slug, "is_delete": 0})
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	// There cannot be an asset in the collection otherwise deleting failed
	if row >= 1 {
		return CONNOT_DELETE_COLLECTION_ERR
	}

	if _, err := db.Collection("collections").UpdateMany(
		context.TODO(),
		bson.M{"userMetamaskId": user, "id": slug, "isDelete": 0},
		bson.M{"$set": bson.M{"isDelete": 1}}); err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
