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
		return
	}

	// Recuperer le contenu du message & le destinataire
	type Body struct {
		Recipient uint
		Message   string
	}
	var mBody Body

	err = c.Bind(&mBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "bad request",
			"error":  err,
		})
		return
	}

	//Enregistrer le message
	m := models.Message{
		UserSender:    u,
		UserRecipient: mBody.Recipient,
		Message:       mBody.Message,
	}
	result := config.DB.Save(&m)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "bad request",
			"error":  result.Error,
		})
		return
	}
	//Retour
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": m,
	})
}
