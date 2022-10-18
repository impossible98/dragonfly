package format

import (
	// import built-in packages
	"fmt"
	// import third-party packages
	"github.com/gin-gonic/gin"
	// import local packages
	"dragonfly/app/utils/ecode"
)

func HTTP(c *gin.Context, code int, err error, data interface{}) {
	type resp struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data,omitempty"`
	}

	var msg = ecode.GetMsg(code)
	if err != nil {
		msg += fmt.Sprintf(": %s", err.Error())
	}
	c.JSON(200, &resp{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
