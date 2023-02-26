package controllers

import (
	"API_go/go_test/config"
	"API_go/go_test/models"
	"net/http"
	"strconv"
	"time"

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
	// Get the recipient id

	if mBody.Recipient == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": "recipient not found",
		})
	}
	var r models.Contact
	config.DB.First(&r, mBody.Recipient)

	//Save the message
	m := models.Message{
		UserSender:    u,
		UserRecipient: r.UserContactID,
		Message:       mBody.Message,
	}
	result := config.DB.Save(&m)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "saving messsage failed",
		})
		return
	}
	//Succes message
	c.JSON(http.StatusCreated, gin.H{
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
	idParam := c.Param("id")
	var r models.Contact

	id, _ := strconv.Atoi(idParam)

	result := config.DB.First(&r, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	//Find all the messages
	type Message struct {
		ID        uint
		CreatedAt time.Time
		Firstname string
		Sender    bool
		Message   string
	}
	var m []Message

	result = config.DB.Raw(`
	SELECT IF(s.id = ?, r.firstname, s.firstname) as Firstname, IF(s.id= ?, true, false) as Sender, m.id, message, m.created_at 
	FROM messages m 
	INNER JOIN users s ON m.user_sender = s.id
	INNER JOIN users r ON m.user_recipient = r.id
	WHERE (s.id = ? AND r.id= ? ) OR (s.id = ? AND r.id = ?)
	ORDER BY m.created_at
	LIMIT 100;
	`, u, u, u, r.UserContactID, r.UserContactID, u).Scan(&m)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
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

func DeleteMessage(c *gin.Context) {
	// Get the message id
	id := c.Param("id")

	//Delete the message
	result := config.DB.Unscoped().Delete(&models.Message{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "bad request",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "ok",
		"message": "message deleted",
	})
}

func GetLastMessage(c *gin.Context) {
	idParam := c.Param("id")
	var r models.Contact

	id, _ := strconv.Atoi(idParam)

	result := config.DB.First(&r, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	type LastMessage struct {
		Message string
	}

	var lastMessage LastMessage

	result = config.DB.Raw(`
	SELECT message FROM contacts c 
	JOIN users u ON c.user_contact_id = u.id 
	JOIN messages m ON m.user_recipient = u.id OR m.user_sender = u.id 
	WHERE c.user_contact_id = ? 
	ORDER BY m.created_at DESC 
	LIMIT 1
	`, r.UserContactID).Scan(&lastMessage)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": lastMessage,
	})
}
