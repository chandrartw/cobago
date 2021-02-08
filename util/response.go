package util

import (
	"github.com/gin-gonic/gin"
)

func ResponseError(ctx *gin.Context, status int, debug string, msg string) {
	result := gin.H{
		"result": msg,
		"debug":  debug,
		"error":  1,
	}

	ctx.AbortWithStatusJSON(status, result)

}

func ResponseSuccess(ctx *gin.Context, status int, structPtr interface{}) {
	result := gin.H{
		"result": structPtr,
		"error":  0,
	}
	ctx.AbortWithStatusJSON(status, result)

}

func ResponseSuccessCustomMessage(ctx *gin.Context, status int, msg string) {

	result := gin.H{
		"result": msg,
		"error":  0,
	}
	ctx.AbortWithStatusJSON(status, result)
}
