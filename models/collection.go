package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"openseasync/database"
	"openseasync/logs"
	"time"
)

var CONNOT_DELETE_COLLECTION_ERR = errors.New("Cannot delete a collection that has an asset")

type Collection struct {
	ID              int64   `json:"id"`                // 主键
	Slug            string  `json:"slug"`              // 集合唯一标识符
	UserAddress     string  `json:"user_address"`      // 集合拥有者
	Name            string  `json:"name"`              // 集合名称
	BannerImageURL  string  `json:"banner_image_url"`  // 集合背景图
	Description     string  `json:"description"`       // 集合描述
	ImageURL        string  `json:"image_url"`         // 集合头像
	LargeImageURL   string  `json:"large_image_url"`   // 头像大图
	IsDelete        int8    `json:"is_delete"`         // 是否删除 1删除 0未删除 默认为0
	CreateDate      string  `json:"create_date"`       // 集合创建时间
	RefreshTime     int     `json:"refresh_time"`      // 刷新时间
	NumOwners       int     `json:"num_owners"`        // 集合中输入自己的NFT个数
	TotalSupply     int     `json:"total_supply"`      // 集合中NFT总数
	TotalVolume     float64 `json:"total_volume"`      // 交易量
	OwnedAssetCount string  `json:"owned_asset_count"` // 所有NFT中属于自己的NFT个数 此地段可能是个big int, 所以采用string存储
}

// InsertOpenSeaCollection find collection through opensea API and insert
func InsertOpenSeaCollection(collections *OwnerCollection, user string) error {
	db := database.GetDB()
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
			OwnedAssetCount: v.OwnedAssetCount.String(),
		}
		// gorm v1  batch insert is not supported
		var tmp Collection
		rows := db.Table("collections").
			Where("user_address = ? AND slug = ? AND is_delete = 0", user, v.Slug).
			Find(&tmp).RowsAffected
		if rows == 0 {
			// insert
			if err := db.Table("collections").
				Create(&collection).Error; err != nil && err != gorm.ErrRecordNotFound {
				logs.GetLogger().Error(err)
				return err
			}
		} else {
			// update
			fmt.Println(user, v.Slug)
			if err := db.Table("collections").Where("user_address = ? AND slug = ? AND is_delete = 0", user, v.Slug).
				Updates(map[string]interface{}{
					"name":             v.Name,
					"description":      v.Description,
					"banner_image_url": v.BannerImageURL,
					"image_url":        v.ImageURL,
					"large_image_url":  v.LargeImageURL,
					"refresh_time":     refreshTime,
				}).Error; err != nil && err != gorm.ErrRecordNotFound {
				logs.GetLogger().Error(err)
				return err
			}
		}

	}

	if err := db.Table("collections").
		Where("user_address = ? AND refresh_time < ? AND is_delete = 0", user, refreshTime).
		Update("is_delete", 1).Error; err != nil && err != gorm.ErrRecordNotFound {
		logs.GetLogger().Error(err)
		return err
	}
	return nil

}

// FindCollectionByOwner find collections by owner
func FindCollectionByOwner(user string) ([]*Collection, error) {
	var collections []*Collection
	db := database.GetDB()
	if err := db.Table("collections").
		Where("user_address = ? AND is_delete = 0", user).
		Find(&collections).Error; err != nil && err != gorm.ErrRecordNotFound {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return collections, nil
}

// DeleteCollectionBySlug delete empty collection
func DeleteCollectionBySlug(user, slug string) error {
	var row int
	db := database.GetDB()
	if err := db.Table("assets").
		Where("user_address = ? AND slug = ? AND is_delete = 0", user, slug).
		Count(&row).Error; err != nil && err != gorm.ErrRecordNotFound {
		logs.GetLogger().Error(err)
		return err
	}
	// There cannot be an asset in the collection otherwise deleting failed
	if row >= 1 {
		return CONNOT_DELETE_COLLECTION_ERR
	}

	if err := db.Table("collections").
		Where("user_address = ? AND is_delete= 0 AND slug = ?", user, slug).
		Update("is_delete", 1).Error; err != nil && err != gorm.ErrRecordNotFound {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
