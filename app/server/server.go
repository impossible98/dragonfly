package server

import (
	// import built-in packages
	"strconv"
	// import third-party packages
	"github.com/gin-gonic/gin"
	// import local packages
	"dragonfly/app/api"
	"dragonfly/app/api/v1"
	"dragonfly/app/config"
)

func SetupServer() {
	app := gin.Default()
	app.GET("/api/version", api.ReadVersion)
	app.GET("/api/healthcheck", api.ReadHealthcheck)
	apiV1 := app.Group("/api/v1")
	{
		apiV1.POST("/favourite", v1.CreateFavourite)
	}
	app.Run(":" + strconv.Itoa(config.ServerPort))
}
