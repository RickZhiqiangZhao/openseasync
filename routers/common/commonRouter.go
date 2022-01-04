package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"openseasync/common"
	"openseasync/common/constants"
	"openseasync/common/errorinfo"
	"openseasync/logs"
	"openseasync/models"
)

func HostManager(router *gin.RouterGroup) {
	router.GET(constants.URL_HOST_GET_HOST_INFO, GetSwanMinerVersion)
	router.GET(constants.URL_OPENSEA_OWNER_ASSETS_SYNC, OpenSeaOwnerDataSync)
	//router.GET(constants.URL_OPENSEA_OWNER_ASSETS, OpenSeaOwnerAssetsSync)
	//router.GET(constants.URL_OPENSEA_OWNER_Collections, OpenSeaOwnerCollectionsSync)
	router.GET(constants.URL_FIND_ASSET, GetAssetsByOwner)
	router.GET(constants.URL_FIND_COLLECTION, GetCollectionsByOwner)
	router.GET(constants.URL_FIND_ASSETS_SLUG, GetAssetsBySlug)
	router.DELETE(constants.URL_DELETE_ASSET, DeleteAssetByTokenID)
	router.DELETE(constants.URL_DELETE_COLLECTION, DeleteCollectionBySlug)

}

func GetSwanMinerVersion(c *gin.Context) {
	info := getSwanMinerHostInfo()
	c.JSON(http.StatusOK, common.CreateSuccessResponse(info))
}

// sync opensea assets and collections
func OpenSeaOwnerDataSync(c *gin.Context) {
	user := c.Param("user")
	if len(user) != 42 {
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_MSG))
		return
	}

	// sync assets
	if err := openSeaOwnerAssetsSync(user); err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.OPENSEA_HTTP_REQUEST_ERROR_CODE, err.Error()))
		return
	}
	// sync collections
	if err := openSeaOwnerCollectionsSync(user); err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.OPENSEA_HTTP_REQUEST_ERROR_CODE, err.Error()))
		return
	}
	c.JSON(http.StatusOK, common.CreateSuccessResponse(nil))
}

// Deprecated: Recommended use OpenSeaOwnerDataSync
func OpenSeaOwnerAssetsSync(c *gin.Context) {
	user := c.Param("user")
	if len(user) != 42 {
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_MSG))
		return
	}
	if err := openSeaOwnerAssetsSync(user); err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.OPENSEA_HTTP_REQUEST_ERROR_CODE, err.Error()))
		return
	}
	c.JSON(http.StatusOK, common.CreateSuccessResponse(nil))
}

// Deprecated: Recommended use OpenSeaOwnerDataSync
func OpenSeaOwnerCollectionsSync(c *gin.Context) {
	user := c.Param("user")
	if len(user) != 42 {
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_MSG))
		return
	}
	if err := openSeaOwnerCollectionsSync(user); err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.OPENSEA_HTTP_REQUEST_ERROR_CODE, err.Error()))
		return
	}
	c.JSON(http.StatusOK, common.CreateSuccessResponse(nil))
}

func GetAssetsByOwner(c *gin.Context) {
	user := c.Param("user")
	if len(user) != 42 {
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_MSG))
		return
	}
	assets, err := getAssetByOwner(user)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_RECORD_lIST_ERROR_CODE, err.Error()))
		return
	}
	c.JSON(http.StatusOK, common.CreateSuccessResponse(assets))
}

func GetCollectionsByOwner(c *gin.Context) {
	user := c.Param("user")
	if len(user) != 42 {
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_MSG))
		return
	}
	collections, err := getCollectionsByOwner(user)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_RECORD_lIST_ERROR_CODE, err.Error()))
		return
	}
	c.JSON(http.StatusOK, common.CreateSuccessResponse(collections))
}

func GetAssetsBySlug(c *gin.Context) {
	user := c.Param("user")
	slug := c.Param("slug")
	if len(user) != 42 && slug != "" {
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_MSG))
		return
	}
	assets, err := getAssetBySlug(user, slug)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_RECORD_lIST_ERROR_CODE, err.Error()))
		return
	}
	c.JSON(http.StatusOK, common.CreateSuccessResponse(assets))
}

func DeleteAssetByTokenID(c *gin.Context) {
	user := c.Param("user")
	contractAddress := c.Param("contract_address")
	tokenID := c.Param("token_id")

	if len(user) != 42 && len(contractAddress) != 42 && tokenID != "" {
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_MSG))
		return
	}
	err := deleteAssetByTokenID(user, contractAddress, tokenID)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.UPDATE_DATA_TO_DB_ERROR_CODE, err.Error()))
		return
	}
	c.JSON(http.StatusOK, common.CreateSuccessResponse(nil))
}

func DeleteCollectionBySlug(c *gin.Context) {
	user := c.Param("user")
	slug := c.Param("slug")
	if len(user) != 42 && slug != "" {
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_MSG))
		return
	}
	err := deleteCollectionBySlug(user, slug)
	if err == models.CONNOT_DELETE_COLLECTION_ERR {
		c.JSON(http.StatusOK, common.CreateErrorResponse(errorinfo.UPDATE_DATA_TO_DB_ERROR_CODE, err.Error()))
		return
	} else if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_RECORD_lIST_ERROR_CODE, err.Error()))
		return
	}
	c.JSON(http.StatusOK, common.CreateSuccessResponse(nil))
}
