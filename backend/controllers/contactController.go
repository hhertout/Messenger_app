package controllers

import (
	"API_go/go_test/config"
	"API_go/go_test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetContacts(c *gin.Context) {
	// Get the user connected
	u, _ := c.Get("userid")

	// Get all contacts of the user
	var allContacts []models.Contact

	result := config.DB.Where("user_id = ?", u).Find(&allContacts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error,
		})
		return
	}

	type Contact struct {
		ID        uint
		Email     string
		Firstname string
		Lastname  string
	}

	// Expose the result => contacts
	/* 	c.JSON(http.StatusOK, gin.H{
		"contacts_nb": result.RowsAffected,
		"contacts":    contacts,
	}) */
}

func ShowContact(c *gin.Context) {
	id := c.Param("id")
	var contact models.Contact
	config.DB.First(&contact, id)

	c.JSON(http.StatusOK, gin.H{
		"contact": contact,
	})
}
