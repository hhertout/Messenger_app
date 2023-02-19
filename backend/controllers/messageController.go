package controllers

import (
	"API_go/go_test/config"
	"API_go/go_test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendMessage(c *gin.Context) {
	// Recuperer le user connecter
	u, err := getUserConnected(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "undefined user",
			"err":     err,
		})
	}

	// Recuperer le user destinataire
	recipientId := c.Param("id")
	type User struct {
		ID uint
	}
	var r User
	config.DB.First(&r, recipientId)

	// Recuperer le contenu du message
	type Body struct {
		Message string
	}
	var mBody Body
	c.Bind(&mBody)

	//Enregistrer le message
	m := models.Message{
		Message:       mBody.Message,
		UserSender:    u,
		UserRecipient: r.ID,
	}
	config.DB.Save(&m)
	//Retour
}
