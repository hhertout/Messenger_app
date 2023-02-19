package controllers

import (
	"API_go/go_test/config"
	"API_go/go_test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var statusAccepted, statusPending = "accepted", "pending"

func Invite(c *gin.Context) {
	// Recuperer l'ID du user connecté
	userId, err := getUserConnected(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "user log error",
		})
	}

	// Recuperer le user a qui l'invitation est envoyée
	type Body struct {
		ID uint
	}
	var iBody Body
	if c.Bind(&iBody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "bad request",
			"message": "invitation failed",
		})
	}

	// Enregistré l'invitation avec le status pending
	invitation := models.Invitation{
		UserRecipient: iBody.ID,
		UserSendeur:   userId,
		Status:        statusPending,
	}

	config.DB.Save(&invitation)

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "invitation successfully sent",
	})
}

// Modification du status de l'invitation

// Suppresssion de l'invitation
func DeleteInvitation(c *gin.Context) {
	id := c.Param("id")

	result := config.DB.Unscoped().Delete(&models.Invitation{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "invitation's delete failed",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "invitation successfully deleted",
	})
}

// Toutes les invitations
func GetInvitation(c *gin.Context) {
	// Recuperer l'ID du user connecté
	userId, err := getUserConnected(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "user log error",
		})
	}

	// Recuperer les invitations du user connecté
	type Users struct {
		Firstname string
		Lastname  string
		Status    string
	}

	var invitations []Users

	result := config.DB.Raw(`
	SELECT firstname, lastname, status 
	FROM invitations i 
	INNER JOIN users u on i.user_sendeur = u.id 
	WHERE i.status = ? 
	AND i.user_sendeur = ? 
	ORDER BY i.created_at
	`, statusPending, userId).Scan(&invitations)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot retrieve data",
			"error":   result.Error,
		})
	}

	// Envoyer les invitations
	c.JSON(http.StatusOK, gin.H{
		"invitationsNumber": result.RowsAffected,
		"invitations":       invitations,
	})
}

func GetAcceptedInvitation(c *gin.Context) {
	// Recuperer l'ID du user connecté
	userId, err := getUserConnected(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "user log error",
		})
	}

	// Recuperer les invitations du user connecté
	type Users struct {
		Firstname string
		Lastname  string
		Status    string
	}

	var invitations []Users

	result := config.DB.Raw(`
	SELECT firstname, lastname, status 
	FROM invitations i 
	INNER JOIN users u on i.user_sendeur = u.id 
	WHERE i.status = ? 
	AND i.user_sendeur = ? 
	ORDER BY i.updated_at
	`, statusAccepted, userId).Scan(&invitations)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot retrieve data",
			"error":   result.Error,
		})
	}

	// Envoyer les invitations
	c.JSON(http.StatusOK, gin.H{
		"invitationsNumber": result.RowsAffected,
		"invitations":       invitations,
	})
}
