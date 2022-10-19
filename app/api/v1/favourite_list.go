package v1

import (
	// import third-party packages
	"github.com/gin-gonic/gin"
	// import local packages
	"dragonfly/app/service/favourite"
	"dragonfly/app/utils/ecode"
	"dragonfly/app/utils/format"
)

func ReadFavouriteList(ctx *gin.Context) {
	result, err := favourite.ReadFavouriteList()
	// control flow
	if err != nil {
		format.HTTP(ctx, ecode.ErrorReadFavouriteList, err, nil)
		// return
		return
	} else {
		format.HTTP(ctx, ecode.Success, nil, result)

	}
}
