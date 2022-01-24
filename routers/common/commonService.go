package common

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"openseasync/common"
	"openseasync/common/utils"
	"openseasync/logs"
	"openseasync/models"
	"runtime"
	"time"
)

func getSwanMinerHostInfo() *common.HostInfo {
	info := new(common.HostInfo)
	info.SwanMinerVersion = common.GetVersion()
	info.OperatingSystem = runtime.GOOS
	info.Architecture = runtime.GOARCH
	info.CPUnNumber = runtime.NumCPU()
	return info
}

// openSeaOwnerAssetsSync get all assets by owner
func openSeaOwnerAssetsSync(user string) error {
	var n int64 = 1
	refreshTime := time.Now().Unix()
	for {
		// If the number of requests is too many, a 429 error code will be thrown
		content, err := utils.RequestOpenSeaAssets(user, 50*(n-1), 50)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
		var assets models.OwnerAsset
		if err = json.Unmarshal(content, &assets); err != nil {
			logs.GetLogger().Error(err)
			return err
		}
		if len(assets.Assets) < 1 {
			break
		}
		if err = models.InsertOpenSeaAsset(&assets, user, refreshTime); err != nil {
			logs.GetLogger().Error(err)
			return err
		}
		if len(assets.Assets) < 50 {
			break
		}
		n++
		//time.Sleep(time.Second)
	}

	return nil
}

// openSeaOwnerCollectionsSync get all collections by owner
func openSeaOwnerCollectionsSync(user string) error {
	var n int64 = 1
	refreshTime := time.Now().Unix()
	for {
		//time.Sleep(time.Second * 2)
		content, err := utils.RequestOpenSeaCollections(user, 300*(n-1), 300*n)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
		var collections models.OwnerCollection
		if err = json.Unmarshal(content, &collections.Collections); err != nil {
			logs.GetLogger().Error(err)
			return err
		}
		if err = models.InsertOpenSeaCollection(&collections, user, refreshTime); err != nil {
			return err
		}
		if len(collections.Collections) < 300 {
			break
		}
		n++
	}

	return nil
}

// getAssetSearchByOwner get assets by owner
func getAssetSearchByOwner(collectionId string, param models.Params) (map[string]interface{}, error) {
	result, err := models.FindAssetSearchByOwner(collectionId, param)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return result, nil
}

// getAssetGeneralInfoByCollectibleId get assets by collectibleId
func getAssetGeneralInfoByCollectibleId(collectibleId int64) (map[string]interface{}, error) {
	result, err := models.FindAssetByGeneralInfoCollectibleId(collectibleId)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return result, nil
}

// getAssetOtherByCollection get assets other by collectibleId
func getAssetOtherByCollection(collectibleId int64) (map[string]interface{}, error) {
	result, err := models.FindAssetOtherByCollection(collectibleId)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return result, nil
}

// getCollectionsByOwner get collection by owner
func getCollectionsByUserMetamaskID(usermetamaskid string, page, pageSize int64) (map[string]interface{}, error) {
	result, err := models.FindCollectionByUserMetamaskID(usermetamaskid, page, pageSize)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return result, nil
}

// getCollectionsByCollectionID get collection by slug
func getCollectionsByCollectionID(collectionId string, page, pageSize int64) (map[string]interface{}, error) {
	result, err := models.FindCollectionByCollectionID(collectionId, page, pageSize)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return result, nil
}

// getAssetOfferRecordsByCollectibleId get asset orders by collectibleId
func getAssetOfferRecordsByCollectibleId(collectibleId int64) ([]bson.M, error) {
	result, err := models.FindAssetOfferRecordsByCollectibleId(collectibleId)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return result, nil
}

// getUserMediaByUserId find user media by userId
func getUserMediaByUserId(userId string) (interface{}, error) {
	result, err := models.FindUserMediaByUserId(userId)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return result, nil
}

// getItemActivityByCollectionId get item_activity by collectionId
func getItemActivityByCollectionId(collectionId string, page, pageSize int64) (map[string]interface{}, error) {
	result, err := models.FindItemActivityByCollectionId(collectionId, page, pageSize)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return result, nil
}

// getTradeHistoryByCollectibleId get trade history by collectibleId
func getTradeHistoryByCollectibleId(collectibleId int64, page, pageSize int64) (map[string]interface{}, error) {
	result, err := models.FindTradeHistoryByCollectibleId(collectibleId, page, pageSize)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return result, nil
}

// deleteAssetByTokenID delete asset
func deleteAssetByTokenID(user, contractAddress, tokenID string) error {
	if err := models.DeleteAssetByTokenID(user, contractAddress, tokenID); err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	return nil
}

// deleteCollectionByCollectionId delete collection
func deleteCollectionByCollectionId(user, slug string) error {
	if err := models.DeleteCollectionByCollectionId(user, slug); err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	return nil
}
