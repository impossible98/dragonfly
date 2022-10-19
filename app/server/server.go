package server

import (
	// import built-in packages
	"strconv"
	// import third-party packages
	"github.com/gin-gonic/gin"
	// import local packages
	"dragonfly/app/api"
	"dragonfly/app/config"
)

func SetupServer() {
	app := gin.Default()
	app.GET("/api/version", api.ReadVersion)
	app.GET("/api/healthcheck", api.ReadHealthcheck)
	app.Run(":" + strconv.Itoa(config.ServerPort))
}
