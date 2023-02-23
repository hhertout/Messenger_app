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
	type Contact struct {
		Firstname string
		Lastname  string
	}
	var allContacts []Contact

	result := config.DB.Raw(`
	SELECT c.ID, firstname, lastname FROM contacts c 
	JOIN users u ON c.user_contact_id = u.id 
	WHERE c.user_owner_id = ?
	`, u).Scan(&allContacts)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Bad request",
		})
		return
	}

	// Expose the result => contacts
	c.JSON(http.StatusOK, gin.H{
		"contacts_nb": result.RowsAffected,
		"contacts":    allContacts,
	})
}

func addContact(UserSender, UserReceiver uint) (models.Contact, error) {
	//Recuperer les deux userid => en args
	//Créer le contact
	c1 := models.Contact{
		UserOwnerID:   UserSender,
		UserContactID: UserReceiver,
	}

	result := config.DB.Save(&c1)
	if result.Error != nil {
		return c1, result.Error
	}

	c2 := models.Contact{
		UserOwnerID:   UserReceiver,
		UserContactID: UserSender,
	}
	result = config.DB.Save(&c2)
	if result.Error != nil {
		return c1, result.Error
	}

	//Renvoyer le contact créer
	return c1, nil
}
