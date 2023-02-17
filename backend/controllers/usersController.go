package controllers

import (
	"API_go/go_test/config"
	"API_go/go_test/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Body struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func Signup(ctx *gin.Context) {
	body := Body{}

	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Hash Failed",
		})
	}

	user := models.User{
		Email:    body.Email,
		Password: string(hash),
	}
	result := config.DB.Create(&user)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "User creation failed",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": "User created",
	})
}
