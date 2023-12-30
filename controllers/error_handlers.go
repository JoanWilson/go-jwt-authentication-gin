package controllers

import "github.com/gin-gonic/gin"

func ErrorJSONMaker(e error) gin.H {
	return gin.H{
		"error": e.Error(),
	}
}

func ErrorContextHandler(ctx *gin.Context, code int, err error) {
	ctx.JSON(code, ErrorJSONMaker(err))
	ctx.Abort()
}
