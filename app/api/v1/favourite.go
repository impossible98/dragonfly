package v1

import (
	// import third-party packages
	"github.com/gin-gonic/gin"
	// import local packages
	"dragonfly/app/service/favourite"
	"dragonfly/app/utils/ecode"
	"dragonfly/app/utils/format"
)

type Request struct {
	Platform string `json:"platform" binding:"required"`
	RoomId   string `json:"room_id" binding:"required" `
}

func CreateFavourite(ctx *gin.Context) {
	request := Request{}
	// control flow
	if err := ctx.ShouldBind(&request); err != nil {
		format.HTTP(ctx, ecode.InvalidParams, err, nil)
		// return
		return
	}
	result, err := favourite.CreateFavourite(request.Platform, request.RoomId)
	// control flow
	if err != nil {
		format.HTTP(ctx, ecode.ErrorCreateFavourite, err, nil)
		// return
		return
	}
	format.HTTP(ctx, ecode.Success, nil, result)
}

func ReadFavourite(ctx *gin.Context) {
	platform := ctx.Query("platform")
	roomId := ctx.Query("room_id")
	result, err := favourite.ReadFavourite(platform, roomId)
	// control flow
	if err != nil {
		format.HTTP(ctx, ecode.ErrorReadFavourite, err, nil)
		// return
		return
	}
	format.HTTP(ctx, ecode.Success, nil, result)
}

func UpdateFavourite(ctx *gin.Context) {
	request := Request{}
	// control flow
	if err := ctx.ShouldBind(&request); err != nil {
		format.HTTP(ctx, ecode.InvalidParams, err, nil)
		// return
		return
	}
	result, err := favourite.UpdateFavourite(request.Platform, request.RoomId)
	// control flow
	if err != nil {
		format.HTTP(ctx, ecode.ErrorUpdateFavourite, err, nil)
		// return
		return
	}
	format.HTTP(ctx, ecode.Success, nil, result)
}

func DeleteFavourite(ctx *gin.Context) {
	request := Request{}
	// control flow
	if err := ctx.ShouldBind(&request); err != nil {
		format.HTTP(ctx, ecode.InvalidParams, err, nil)
		// return
		return
	}
	result, err := favourite.DeleteFavourite(request.Platform, request.RoomId)
	// control flow
	if err != nil {
		format.HTTP(ctx, ecode.ErrorUpdateFavourite, err, nil)
		// return
		return
	}
	format.HTTP(ctx, ecode.Success, nil, result)
}
