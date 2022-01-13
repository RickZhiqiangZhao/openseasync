package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"openseasync/common"
	"openseasync/common/constants"
	"openseasync/common/errorinfo"
	"openseasync/logs"
	"openseasync/models"
	"strconv"
)

func HostManager(router *gin.RouterGroup) {
	router.GET(constants.URL_HOST_GET_HOST_INFO, GetSwanMinerVersion)
	router.GET(constants.URL_OPENSEA_OWNER_ASSETS_SYNC, OpenSeaOwnerDataSync)
	//router.GET(constants.URL_OPENSEA_OWNER_ASSETS, OpenSeaOwnerAssetsSync)
	//router.GET(constants.URL_OPENSEA_OWNER_Collections, OpenSeaOwnerCollectionsSync)
	router.GET(constants.URL_FIND_ASSET, GetAssetsByOwner)
	router.GET(constants.URL_FIND_COLLECTION_USERMETAMASKID, GetCollectionsByUserMetamaskID)
	router.GET(constants.URL_FIND_COLLECTION_COLLECTIONID, GetCollectionsByCollectionID)
	router.GET(constants.URL_FIND_COLLECTION_ITEM_ACTIVITY_COLLECTIONID, GetItemActivityByCollectionID)
	router.GET(constants.URL_FIND_ASSETS_SLUG, GetAssetsBySlug)
	router.DELETE(constants.URL_DELETE_ASSET, DeleteAssetByTokenID)
	router.DELETE(constants.URL_DELETE_COLLECTION, DeleteCollectionBySlug)

}

func GetSwanMinerVersion(c *gin.Context) {
	info := getSwanMinerHostInfo()
	c.JSON(http.StatusOK, common.CreateSuccessResponse(info, nil))
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
	c.JSON(http.StatusOK, common.CreateSuccessResponse(nil, nil))
}

func GetAssetsByOwner(c *gin.Context) {
	user := c.Param("user")
	page := c.Query("page")
	pageSize := c.Query("pageSize")
	if len(user) != 42 {
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_MSG))
		return
	}
	pageInt, err := strconv.ParseInt(page, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_MSG))
		return
	}
	pageSizeInt, err := strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_MSG))
		return
	}
	if pageInt < 1 || pageSizeInt < 1 {
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_MSG))
		return
	}

	result, err := getAssetByOwner(user, pageInt, pageSizeInt)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_RECORD_lIST_ERROR_CODE, err.Error()))
		return
	}
	c.JSON(http.StatusOK, common.CreateSuccessResponse(result["data"], result["metadata"]))
}

func GetCollectionsByUserMetamaskID(c *gin.Context) {
	usermetamaskid := c.Param("usermetamaskid")
	page := c.Query("page")
	pageSize := c.Query("pageSize")
	if len(usermetamaskid) != 42 {
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_MSG))
		return
	}
	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_MSG))
		return
	}
	pageSizeInt, err := strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_MSG))
		return
	}
	if pageInt < 1 || pageSizeInt < 1 {
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_MSG))
		return
	}
	result, err := getCollectionsByUserMetamaskID(usermetamaskid, pageInt, pageSizeInt)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_RECORD_lIST_ERROR_CODE, err.Error()))
		return
	}
	c.JSON(http.StatusOK, common.CreateSuccessResponse(result["data"], result["metadata"]))
}

func GetCollectionsByCollectionID(c *gin.Context) {
	collectionId := c.Param("collectionId")
	page := c.Query("page")
	pageSize := c.Query("pageSize")
	if collectionId == "" {
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_MSG))
		return
	}
	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_MSG))
		return
	}
	pageSizeInt, err := strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_MSG))
		return
	}
	if pageInt < 1 || pageSizeInt < 1 {
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_MSG))
		return
	}
	result, err := getCollectionsByCollectionID(collectionId, pageInt, pageSizeInt)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_RECORD_lIST_ERROR_CODE, err.Error()))
		return
	}
	c.JSON(http.StatusOK, common.CreateSuccessResponse(result["data"], result["metadata"]))
}

func GetItemActivityByCollectionID(c *gin.Context) {
	collectionId := c.Param("collectionId")
	page := c.Query("page")
	pageSize := c.Query("pageSize")
	if collectionId == "" {
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_MSG))
		return
	}
	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_MSG))
		return
	}
	pageSizeInt, err := strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_MSG))
		return
	}
	if pageInt < 1 || pageSizeInt < 1 {
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_MSG))
		return
	}
	result, err := getItemActivityByCollectionId(collectionId, pageInt, pageSizeInt)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_RECORD_lIST_ERROR_CODE, err.Error()))
		return
	}
	c.JSON(http.StatusOK, common.CreateSuccessResponse(result["data"], result["metadata"]))
}

func GetAssetsBySlug(c *gin.Context) {
	user := c.Param("user")
	slug := c.Param("slug")
	page := c.Query("page")
	pageSize := c.Query("pageSize")
	if len(user) != 42 && slug != "" {
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_MSG))
		return
	}
	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_MSG))
		return
	}
	pageSizeInt, err := strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_MSG))
		return
	}
	if pageInt < 1 || pageSizeInt < 1 {
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAM_VALUE_ERROR_MSG))
		return
	}
	result, err := getAssetBySlug(user, slug, pageInt, pageSizeInt)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_RECORD_lIST_ERROR_CODE, err.Error()))
		return
	}
	c.JSON(http.StatusOK, common.CreateSuccessResponse(result["data"], result["metadata"]))
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
	c.JSON(http.StatusOK, common.CreateSuccessResponse(nil, nil))
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
	c.JSON(http.StatusOK, common.CreateSuccessResponse(nil, nil))
}
