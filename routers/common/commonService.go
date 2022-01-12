package common

import (
	"encoding/json"
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
		if err = models.InsertOpenSeaAsset(&assets, user); err != nil {
			logs.GetLogger().Error(err)
			return err
		}
		if len(assets.Assets) < 50 {
			break
		}
		n++
		time.Sleep(time.Second)
	}

	return nil
}

// openSeaOwnerCollectionsSync get all collections by owner
func openSeaOwnerCollectionsSync(user string) error {
	var n int64 = 1
	for {
		time.Sleep(time.Second * 2)
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
		if err = models.InsertOpenSeaCollection(&collections, user); err != nil {
			return err
		}
		if len(collections.Collections) < 300 {
			break
		}
		n++
	}

	return nil
}

// getAssetByOwner get assets by owner
func getAssetByOwner(user string, page, pageSize int64) (map[string]interface{}, error) {
	result, err := models.FindAssetByOwner(user, nil, page, pageSize)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return result, nil
}

// getAssetBySlug get assets by owner
func getAssetBySlug(user, slug string, page, pageSize int64) (map[string]interface{}, error) {
	result, err := models.FindAssetByOwner(user, slug, page, pageSize)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return result, nil
}

// getCollectionsByOwner get collection by owner
func getCollectionsByOwner(user string, page, pageSize int64) (map[string]interface{}, error) {
	result, err := models.FindCollectionByOwner(user, page, pageSize)
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

// deleteCollectionBySlug delete collection
func deleteCollectionBySlug(user, slug string) error {
	if err := models.DeleteCollectionBySlug(user, slug); err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	return nil
}
