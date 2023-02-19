package controllers

import (
	"API_go/go_test/config"
	"API_go/go_test/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Body struct {
	Email     string `form:"email" json:"email"`
	Password  string `form:"password" json:"password"`
	Firstname string `form:"password" json:"firstname"`
	Lastname  string `form:"password" json:"lastname"`
}

func Signup(c *gin.Context) {
	body := Body{}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Hash Failed",
		})
	}

	user := models.User{
		Email:     body.Email,
		Password:  string(hash),
		Lastname:  body.Lastname,
		Firstname: body.Firstname,
	}

	result := config.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User creation failed",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "User created",
	})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	type User struct {
		ID        int
		Firstname string
		Lastname  string
		Email     string
	}

	var user User

	result := config.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad request",
			"message": result.Error,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"user":   user,
	})
}
