package controllers

import (
	"API_go/go_test/config"
	"API_go/go_test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendMessage(c *gin.Context) {
	// Get the connected user
	u, err := getUserConnected(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "undefined user",
			"err":     err,
		})
		return
	}

	// Retrieve message content & recipient
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

	if mBody.Recipient == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": "recipient not found",
		})
	}

	//Save the message
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
	//Succes message
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": m,
	})
}

func GetMessages(c *gin.Context) {
	// Get connected user
	u, err := getUserConnected(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "undefined user",
			"err":     err,
		})
		return
	}

	// Get the recipient id
	var r models.User
	id := c.Param("id")

	result := config.DB.Select("ID").First(&r, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   result.Error,
			"message": "bad request",
		})
		return
	}

	//Find all the messages
	var m []models.Message
	result = config.DB.Raw(`
	SELECT message, s.firstname, s.lastname FROM messages m 
	INNER JOIN users s ON m.user_sender = s.id
	INNER JOIN users r ON m.user_recipient = r.id
	WHERE (s.id = ? AND r.id= ? ) OR (s.id = ? AND r.id = ?)
	ORDER BY m.created_at
	LIMIT 100;
	`, u, r.ID, r.ID, u).Scan(&m)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   result.Error,
			"message": "bad request",
		})
		return
	}

	//Response
	c.JSON(http.StatusOK, gin.H{
		"status":   "ok",
		"messages": m,
	})
}
