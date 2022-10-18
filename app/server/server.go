package server

import (
	// import third-party packages
	"github.com/gin-gonic/gin"
	// import local packages
	"dragonfly/app/api"
)

func SetupServer() {
	app := gin.Default()
	app.GET("/api/version", api.ReadVersion)
	app.Run(":8080")
}