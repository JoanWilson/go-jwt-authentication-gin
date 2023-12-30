package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Entry(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"isRunning": true,
	})
}
