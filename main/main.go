package main

import (
	"context"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"openseasync/common/constants"
	"openseasync/config"
	"openseasync/database"
	"openseasync/logs"
	"openseasync/routers/common"
	"time"
)

func main() {
	// init database
	db := database.InitMongo()
	config.InitConfig("")
	defer func() {
		err := db.Disconnect(context.TODO())
		if err != nil {
			logs.GetLogger().Error(err)
		}
	}()

	r := gin.Default()
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	v1 := r.Group("/api/v1")
	common.HostManager(v1.Group(constants.URL_HOST_GET_COMMON))
	err := r.Run(":" + config.GetConfig().Port)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}
}
