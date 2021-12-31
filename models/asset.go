package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"openseasync/common/utils"
	"openseasync/database"
	"openseasync/logs"
	"time"
)

type Asset struct {
	ID                int64  `json:"id"`                  // 主键
	UserAddress       string `json:"user_address"`        // 用户地址
	Title             string `json:"title"`               // NFT作品标题
	ImageURL          string `json:"image_url"`           // NFT作品图片
	ImagePreviewURL   string `json:"image_preview_url"`   // NFT作品原图
	ImageThumbnailURL string `json:"image_thumbnail_url"` // NFT作品缩略图
	Description       string `json:"description"`         // NFT作品描述
	ContractAddress   string `json:"contract_address"`    // 合约地址
	TokenId           string `json:"token_id"`            // NFT token id
	NumSales          int    `json:"num_sales"`           // NFT售卖次数
	Owner             string `json:"owner"`               // NFT拥有者
	OwnerImgURL       string `json:"owner_img_url"`       // 拥有者头像
	Creator           string `json:"creator"`             // NFT创造者
	CreatorImgURL     string `json:"creator_img_url"`     // 创造者头像
	TokenMetadata     string `json:"token_metadata"`      // NFT元数据

	Slug string `json:"slug"` // 集合唯一标识符号

	Contract            Contract             `json:"contract"`
	Collection          Collection           `json:"collection"`
	AssetsTopOwnerships []AssetsTopOwnership `json:"assets_top_ownership"`
	Traits              []Trait              `json:"trait"`
	IsDelete            int8                 `json:"is_delete"` // 是否删除 1删除 0未删除 默认为0
	Date                int                  `json:"date"`      // 刷新时间
}

type Contract struct {
	ID           int64  `json:"id"`            // 主键
	Address      string `json:"address"`       // 合约地址
	ContractType string `json:"contract_type"` // 合约类型 semi-fungible可替代 non-fungible 不可替代
	ContractName string `json:"contract_name"` // 合约名字
	Symbol       string `json:"symbol"`        // 符号
	SchemaName   string `json:"schema_name"`   // 合约类型
	TotalSupply  string `json:"total_supply"`  // 总供应量
	Description  string `json:"description"`   // 合约描述
}

type Trait struct {
	ID              int64  `json:"id"`         // 主键
	ContractAddress string `json:"_"`          // 合约地址
	TokenId         string `json:"_"`          // token id
	TraitType       string `json:"trait_type"` // 特征类型
	Value           string `json:"value"`      // 特征值
	DisplayType     string `json:"display_type"`
	MaxValue        int    `json:"max_value"`
	TraitCount      int    `json:"trait_count"` // 数量
	OrderBy         string `json:"order_by"`
	IsDelete        int8   `json:"is_delete"` // 是否删除 1删除 0未删除 默认为0
	Date            int    `json:"date"`      // 刷新时间
}

type AssetsTopOwnership struct {
	ID              int64  `json:"id"`              // 主键
	ContractAddress string `json:"_"`               // 合约地址
	TokenId         string `json:"_"`               // token id
	Owner           string `json:"owner"`           // 所有者地址
	ProfileImgURL   string `json:"profile_img_url"` // 所有者头像
	Quantity        string `json:"quantity"`        // 数量
	IsDelete        int8   `json:"is_delete"`       // 是否删除 1删除 0未删除 默认为0
	Date            int    `json:"date"`            // 刷新时间
}

// InsertOpenSeaAsset query Aseets through opensea API and insert
func InsertOpenSeaAsset(assets *OwnerAsset, user string) error {
	db := database.GetDB()
	date := int(time.Now().Unix())

	// No blocking query opensea assets_top_ownerships
	go queryAssetsTopOwnerShip(db, assets, date)

	for _, v := range assets.Assets {
		owner := user
		if v.Owner.Address != "0x0000000000000000000000000000000000000000" {
			owner = v.Owner.Address
		}

		var asset = Asset{
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
			Date:              date,
		}

		var contract = Contract{
			Address:      v.AssetContract.Address,
			ContractName: v.AssetContract.Name,
			ContractType: v.AssetContract.AssetContractType,
			Symbol:       v.AssetContract.Symbol,
			SchemaName:   v.AssetContract.SchemaName,
			TotalSupply:  v.AssetContract.TotalSupply,
			Description:  v.AssetContract.Description,
		}

		// gorm v1  batch insert is not supported
		var tmp1 Asset
		rows1 := db.Table("assets").
			Where("contract_address = ? AND token_id = ?", v.AssetContract.Address, v.TokenID).
			Find(&tmp1).RowsAffected
		if rows1 == 0 {
			// insert
			if err := db.Table("assets").Create(&asset).Error; err != nil {
				logs.GetLogger().Error(err)
				return err
			}
		} else {
			// Refresh synchronization time
			if err := db.Table("assets").
				Where("contract_address = ? AND token_id = ?", v.AssetContract.Address, v.TokenID).
				Update("date", date).Error; err != nil {
				logs.GetLogger().Error(err)
				return err
			}
			// update NFT without locked metadata
			if v.TokenMetadata == "" {
				if err := db.Table("assets").
					Where("contract_address = ? AND token_id = ?", v.AssetContract.Address, v.TokenID).
					Updates(map[string]interface{}{
						"title":               v.Name,
						"image_url":           v.ImageURL,
						"image_preview_url":   v.ImagePreviewURL,
						"image_thumbnail_url": v.ImageThumbnailURL,
						"description":         v.Description,
						"date":                date,
					}).Error; err != nil {
					logs.GetLogger().Error(err)
					return err
				}
			}

		}

		var tmp2 Contract
		rows2 := db.Table("contracts").
			Where("address = ?", v.AssetContract.Address).
			Find(&tmp2).RowsAffected
		if rows2 == 0 {
			if err := db.Table("contracts").Create(&contract).Error; err != nil {
				logs.GetLogger().Error(err)
				return err
			}
		}

		// update ro insert traits
		for _, v1 := range v.Traits {
			var trait = Trait{
				ContractAddress: v.AssetContract.Address,
				TokenId:         v.TokenID,
				TraitType:       v1.TraitType,
				Value:           v1.Value,
				DisplayType:     v1.DisplayType,
				MaxValue:        v1.MaxValue,
				TraitCount:      v1.TraitCount,
				OrderBy:         v1.Order,
				Date:            date,
			}
			var tmp3 Trait
			rows3 := db.Table("traits").
				Where("contract_address = ? AND token_id = ? AND trait_type = ? AND value = ?", v.AssetContract.Address, v.TokenID, v1.TraitType, v1.Value).
				Find(&tmp3).RowsAffected
			if rows3 == 0 {
				if err := db.Table("traits").Create(&trait).Error; err != nil {
					logs.GetLogger().Error(err)
					return err
				}
			} else {
				// Refresh synchronization time
				if err := db.Table("traits").
					Where("contract_address = ? AND token_id = ? AND trait_type = ? AND value = ?", v.AssetContract.Address, v.TokenID, v1.TraitType, v1.Value).
					Update("date", date).Error; err != nil {
					logs.GetLogger().Error(err)
					return err
				}
			}
		}
	}

	// Delete opensea deleted asset
	if err := db.Table("assets").
		Where("date < ?", date).
		Update("is_delete", 1).Error; err != nil && err != gorm.ErrRecordNotFound {
		logs.GetLogger().Error(err)
		return err
	}
	// Delete opensea deleted traits
	if err := db.Table("traits").
		Where("date < ?", date).
		Update("is_delete", 1).Error; err != nil && err != gorm.ErrRecordNotFound {
		logs.GetLogger().Error(err)
		return err
	}

	return nil

}

// FindAssetByOwner find assets by owner
func FindAssetByOwner(owner string) ([]*Asset, error) {
	var assets []*Asset
	db := database.GetDB()
	if err := db.Table("assets").
		Where("owner = ? AND is_delete = 0", owner).
		Find(&assets).Error; err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	for _, v := range assets {
		if err := db.Table("collections").
			Where("slug = ?", v.Slug).
			Find(&v.Collection).Error; err != nil && err != gorm.ErrRecordNotFound {
			logs.GetLogger().Error(err)
			return nil, err
		}
		if err := db.Table("contracts").
			Where("address = ?", v.ContractAddress).
			Find(&v.Contract).Error; err != nil && err != gorm.ErrRecordNotFound {
			logs.GetLogger().Error(err)
			return nil, err
		}
		if err := db.Table("assets_top_ownerships").
			Where("contract_address = ? AND token_id = ?", v.ContractAddress, v.TokenId).
			Find(&v.AssetsTopOwnerships).Error; err != nil && err != gorm.ErrRecordNotFound {
			logs.GetLogger().Error(err)
			return nil, err
		}
		if err := db.Table("traits").
			Where("contract_address = ? AND token_id = ?", v.ContractAddress, v.TokenId).
			Find(&v.Traits).Error; err != nil && err != gorm.ErrRecordNotFound {
			logs.GetLogger().Error(err)
			return nil, err
		}
	}

	return assets, nil
}

// FindWorksBySlug find assets by collection
func FindWorksBySlug(owner, slug string) ([]*Asset, error) {
	var assets []*Asset
	db := database.GetDB()

	if err := db.Table("assets").
		Where("owner = ? AND slug = ? AND is_delete = 0", owner, slug).
		Find(&assets).Error; err != nil && err != gorm.ErrRecordNotFound {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return assets, nil
}

// DeleteAssetByTokenID delete asset by tokenId
func DeleteAssetByTokenID(contractAddress, tokenID string) error {
	db := database.GetDB()
	if err := db.Table("assets").
		Where("contract_address = ? AND token_id = ?", contractAddress, tokenID).
		Update("is_delete", 1).Error; err != nil && err != gorm.ErrRecordNotFound {
		logs.GetLogger().Error(err)
		return err
	}
	if err := db.Table("traits").
		Where("contract_address = ? AND token_id = ?", contractAddress, tokenID).
		Update("is_delete", 1).Error; err != nil && err != gorm.ErrRecordNotFound {
		logs.GetLogger().Error(err)
		return err
	}
	if err := db.Table("assets_top_ownerships").
		Where("contract_address = ? AND token_id = ?", contractAddress, tokenID).
		Update("is_delete", 1).Error; err != nil && err != gorm.ErrRecordNotFound {
		logs.GetLogger().Error(err)
		return err
	}
	return nil
}

func queryAssetsTopOwnerShip(db *gorm.DB, assets *OwnerAsset, date int) error {
	for _, v := range assets.Assets {
		time.Sleep(time.Second * time.Duration(1))
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
		for _, a := range autoAsset.TopOwnerships {
			var assetsTopOwnership = AssetsTopOwnership{
				ContractAddress: v.AssetContract.Address,
				TokenId:         v.TokenID,
				Owner:           a.Owner.Address,
				ProfileImgURL:   a.Owner.ProfileImgURL,
				Quantity:        a.Quantity,
				Date:            date,
			}

			var tmp4 AssetsTopOwnership
			rows4 := db.Table("assets_top_ownerships").
				Where("contract_address = ? AND token_id = ?", v.AssetContract.Address, v.TokenID).
				Find(&tmp4).RowsAffected
			if rows4 == 0 {
				if err = db.Table("assets_top_ownerships").Create(&assetsTopOwnership).Error; err != nil {
					logs.GetLogger().Error(err)
					return err
				}
			} else {
				// Refresh synchronization time
				if err = db.Table("assets_top_ownerships").
					Where("contract_address = ? AND token_id = ?", v.AssetContract.Address, v.TokenID).
					Update("date", date).Error; err != nil {
					logs.GetLogger().Error(err)
					return err
				}
			}
		}

	}
	// Delete opensea deleted traits
	if err := db.Table("assets_top_ownerships").
		Where("date < ?", date).
		Update("is_delete", 1).Error; err != nil && err != gorm.ErrRecordNotFound {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
