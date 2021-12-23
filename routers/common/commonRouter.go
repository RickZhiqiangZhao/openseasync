package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"openseasync/common"
	"openseasync/common/constants"
)

func HostManager(router *gin.RouterGroup) {
	router.GET(constants.URL_HOST_GET_HOST_INFO, GetSwanMinerVersion)
}

func GetSwanMinerVersion(c *gin.Context) {
	info := getSwanMinerHostInfo()
	c.JSON(http.StatusOK, common.CreateSuccessResponse(info))
}

