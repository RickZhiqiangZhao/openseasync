package models

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
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
			Slug:            v.Slug,
			UserAddress:     user,
			Name:            v.Name,
			Description:     v.Description,
			BannerImageURL:  v.BannerImageURL,
			ImageURL:        v.ImageURL,
			LargeImageURL:   v.LargeImageURL,
			CreateDate:      v.CreatedDate,
			RefreshTime:     refreshTime,
			NumOwners:       v.Stats.NumOwners,
			TotalSupply:     int(v.Stats.TotalSupply),
			TotalVolume:     v.Stats.TotalVolume,
			FloorPrice:      v.Stats.FloorPrice,
			OwnedAssetCount: v.OwnedAssetCount.String(),
		}
		count, err := db.Collection("collections").
			CountDocuments(context.TODO(), bson.M{"user_address": user, "slug": v.Slug, "is_delete": 0})
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
			if _, err = db.Collection("collections").UpdateOne(
				context.TODO(),
				bson.M{"user_address": user, "slug": v.Slug, "is_delete": 0},
				bson.M{"$set": bson.M{"name": v.Name, "description": v.Description, "banner_image_url": v.BannerImageURL,
					"image_url": v.ImageURL, "large_image_url": v.LargeImageURL, "refresh_time": refreshTime, "floor_price": v.Stats.FloorPrice,
				}}); err != nil {
				logs.GetLogger().Error(err)
				return err
			}
		}
	}

	if _, err := db.Collection("collections").UpdateMany(
		context.TODO(),
		bson.M{"user_address": user, "refresh_time": bson.M{"$lt": refreshTime}, "is_delete": 0},
		bson.M{"$set": bson.M{"is_delete": 1}}); err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	return nil

}

// FindCollectionByOwner find collections by owner
func FindCollectionByOwner(user string) ([]*Collection, error) {
	var collections []*Collection
	db := database.GetMongoClient()
	cursor, err := db.Collection("collections").Find(context.TODO(), bson.M{"user_address": user, "is_delete": 0})
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	if err = cursor.All(context.TODO(), &collections); err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return collections, nil
}

// DeleteCollectionBySlug delete empty collection
func DeleteCollectionBySlug(user, slug string) error {
	db := database.GetMongoClient()
	row, err := db.Collection("assets").CountDocuments(
		context.TODO(), bson.M{"user_address": user, "slug": slug, "is_delete": 0})
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
		bson.M{"user_address": user, "slug": slug, "is_delete": 0},
		bson.M{"$set": bson.M{"is_delete": 1}}); err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
