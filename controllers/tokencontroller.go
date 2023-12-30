package controllers

import (
	"github.com/gin-gonic/gin"
	"jwt-authentication-golang/auth"
	"jwt-authentication-golang/database"
	"jwt-authentication-golang/models"
	"net/http"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func TokenBuilder(ctx *gin.Context) {
	var req TokenRequest
	var user models.User

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorContextHandler(ctx, http.StatusBadRequest, err)
		return
	}

	record := database.Instance.Where("email = ?", req.Email).First(&user)
	if record.Error != nil {
		ErrorContextHandler(ctx, http.StatusInternalServerError, record.Error)
		return
	}

	credentialError := user.PasswordChecker(req.Password)
	if credentialError != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		ctx.Abort()
		return
	}

	tokenString, err := auth.JWTBuilder(user.Email, user.Username)
	if err != nil {
		ErrorContextHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}
