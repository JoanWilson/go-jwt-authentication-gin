package controllers

import (
	"github.com/gin-gonic/gin"
	"jwt-authentication-golang/database"
	"jwt-authentication-golang/models"
	"net/http"
)

func UserRegister(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ErrorContextHandler(ctx, http.StatusBadRequest, err)
		return
	}

	if err := user.PasswordHasher(user.Password); err != nil {
		ErrorContextHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	record := database.Instance.Create(&user)
	if record.Error != nil {
		ErrorContextHandler(ctx, http.StatusInternalServerError, record.Error)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"userId":   user.ID,
		"email":    user.Email,
		"username": user.Username,
	})
}
