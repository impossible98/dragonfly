package api

import (
	// import third-party packages
	"github.com/gin-gonic/gin"
	// import local packages
	"dragonfly/app/utils/ecode"
	"dragonfly/app/utils/format"
	"dragonfly/global"
)

func ReadVersion(ctx *gin.Context) {
	format.HTTP(ctx, ecode.Success, nil, gin.H{
		"version": global.VERSION,
	})
}
