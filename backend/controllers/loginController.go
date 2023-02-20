package controllers

import (
	"API_go/go_test/config"
	"API_go/go_test/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func Login(c *gin.Context) {
	var body Body

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})

		return
	}

	var user models.User
	config.DB.Where("email= ?", body.Email).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password or email",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password or email",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   os.Getenv("SECRET"),
			"message": "Failed to create token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "/", "http://localhost", true, true)
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "successfully logged",
	})
}

// Recuperer l'ID du user connecté
func getUserConnected(c *gin.Context) (uint, error) {
	u, exist := c.Get("userid")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "bad request",
			"message": "user does not exist",
		})
	}

	var user models.User
	result := config.DB.Select("ID").First(&user, u)
	if result.Error != nil {
		return 0, result.Error
	}
	return uint(user.ID), nil
}
