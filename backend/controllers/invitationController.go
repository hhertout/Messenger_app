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
			"message": "user login error",
		})
		return
	}

	// Recuperer le user a qui l'invitation est envoyée
	type Body struct {
		Email string
	}
	var iBody Body
	if c.Bind(&iBody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "bad request",
			"message": "bad request",
		})
		return
	}
	var r models.User
	result := config.DB.Where(models.User{Email: iBody.Email}).First(&r)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "user not found",
		})
		return
	}

	// Enregistré l'invitation avec le status pending
	invitation := models.Invitation{
		UserRecipient: r.ID,
		UserSendeur:   u,
		Status:        statusPending,
	}

	config.DB.Save(&invitation)

	c.JSON(http.StatusCreated, gin.H{
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
			})
			return
		}
		// Invitation Status Accepted = > Ajout du contact
		contact, err := addContact(invit.UserSendeur, invit.UserRecipient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "contact add failed",
			})
			return
		}
		// Reponse => all's ok
		c.JSON(http.StatusCreated, gin.H{
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
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "ok",
		"message": "invitation successfully deleted",
	})
}

// Toutes les invitations
func GetInvitation(c *gin.Context) {
	// Recuperer l'ID du user connecté
	u, err := getUserConnected(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user log error",
		})
	}

	// Recuperer les invitations du user connecté
	type Users struct {
		ID        uint
		Firstname string
		Lastname  string
		Status    string
	}

	var invitations []Users

	result := config.DB.Raw(`
	SELECT i.id, firstname, lastname, status 
	FROM invitations i 
	INNER JOIN users u on i.user_recipient = u.id 
	WHERE i.status = ? 
	AND i.user_recipient = ? 
	ORDER BY i.created_at
	`, statusPending, u).Scan(&invitations)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot retrieve data",
		})
		return
	}

	// Envoyer les invitations
	c.JSON(http.StatusCreated, gin.H{
		"invitationsNumber": result.RowsAffected,
		"invitations":       invitations,
	})
}

func GetAcceptedInvitation(c *gin.Context) {
	// Recuperer l'ID du user connecté
	u, err := getUserConnected(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "user login error",
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
		})
		return
	}

	// Envoyer les invitations
	c.JSON(http.StatusOK, gin.H{
		"invitationsNumber": result.RowsAffected,
		"invitations":       invitations,
	})
}
