package api

import (
	// import third-party packages
	"github.com/gin-gonic/gin"
)

func ReadHealthcheck(ctx *gin.Context) {
	ctx.String(200, "ok")
}
