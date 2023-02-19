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
	u, err := getUserConnected(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "user log error",
			"error":   err,
		})
		return
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
		return
	}

	// Enregistré l'invitation avec le status pending
	invitation := models.Invitation{
		UserRecipient: iBody.ID,
		UserSendeur:   u,
		Status:        statusPending,
	}

	config.DB.Save(&invitation)

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "invitation sent",
	})
}

// Modification du status de l'invitation
func UpdateInvitation(c *gin.Context) {
	// Recuperer l'id de l'invitation
	id := c.Param("id")

	//Recupération de l'invitation par les paramètres
	var invit models.Invitation
	result := config.DB.First(&invit, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invitation does not exist",
		})
		return
	}

	// Recuprer les changements apporté par l'utilisateur
	type Body struct {
		Status string
	}
	var iBody Body
	if c.Bind(&iBody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "bad request",
			"message": "update failed",
		})
		return
	}
	// Mettre a jour le status de l'invitation
	if iBody.Status == "accepted" {
		//Update du status de l'invitation
		invit.Status = statusAccepted
		result := config.DB.Save(&invit)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "update failed",
				"error":   result.Error,
			})
			return
		}
		// Invitation Status Accepted = > Ajout du contact
		contact, err := addContact(invit.UserSendeur, invit.UserRecipient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "contact add failed",
				"error":   err,
			})
			return
		}
		// Reponse => all's ok
		c.JSON(http.StatusOK, gin.H{
			"status":            "ok",
			"message":           "update done, contact successfully created",
			"contact":           contact,
			"invitation-status": iBody.Status,
		})

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "bad request",
			"message": "update failed",
		})
	}
}

// Suppresssion de l'invitation
func DeleteInvitation(c *gin.Context) {
	id := c.Param("id")

	result := config.DB.Unscoped().Delete(&models.Invitation{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "invitation's delete failed",
			"error":   result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "invitation successfully deleted",
	})
}

// Toutes les invitations
func GetInvitation(c *gin.Context) {
	// Recuperer l'ID du user connecté
	u, err := getUserConnected(c)
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
	`, statusPending, u).Scan(&invitations)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot retrieve data",
			"error":   result.Error,
		})
		return
	}

	// Envoyer les invitations
	c.JSON(http.StatusOK, gin.H{
		"invitationsNumber": result.RowsAffected,
		"invitations":       invitations,
	})
}

func GetAcceptedInvitation(c *gin.Context) {
	// Recuperer l'ID du user connecté
	u, err := getUserConnected(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "user log error",
			"error":   err,
		})
		return
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
	`, statusAccepted, u).Scan(&invitations)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot retrieve data",
			"error":   result.Error,
		})
		return
	}

	// Envoyer les invitations
	c.JSON(http.StatusOK, gin.H{
		"invitationsNumber": result.RowsAffected,
		"invitations":       invitations,
	})
}
