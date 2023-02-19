package controllers

import (
	"API_go/go_test/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetContacts(c *gin.Context) {
	// Get the user connected
	u, _ := c.Get("userid")

	// Get all contacts of the user
	type Contact struct {
		Firstname string
		Lastname  string
	}
	var allContacts []Contact

	result := config.DB.Raw(`
	SELECT firstname, lastname FROM contacts c 
	JOIN users u ON c.user_contact_id = u.id 
	WHERE c.user_owner_id = ?
	`, u).Scan(&allContacts)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Bad request",
			"error":  result.Error,
		})
		return
	}

	// Expose the result => contacts
	c.JSON(http.StatusOK, gin.H{
		"contacts_nb": result.RowsAffected,
		"contacts":    allContacts,
	})
}
