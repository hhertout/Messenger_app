package router

import (
	"API_go/go_test/config"
	"API_go/go_test/controllers"
	"API_go/go_test/middleware"

	"github.com/gin-gonic/gin"
)

func Routes() {
	router := gin.New()
	config.CORSConfig(router)
	// Auth
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.POST("/logout", middleware.RequireAuth, controllers.Logout)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)

	// User
	router.GET("/user/:id", middleware.RequireAuth, controllers.GetUser)

	// Contacts
	router.GET("/contacts", middleware.RequireAuth, controllers.GetContacts)

	// Invitations
	router.GET("/invite/pending", middleware.RequireAuth, controllers.GetInvitation)
	router.GET("/invite/accepted", middleware.RequireAuth, controllers.GetInvitation)
	router.POST("/invite", middleware.RequireAuth, controllers.Invite)
	router.DELETE("/invite/:id", middleware.RequireAuth, controllers.DeleteInvitation)
	router.PUT("/invite/:id", middleware.RequireAuth, controllers.UpdateInvitation)

	// Message
	router.POST("/message", middleware.RequireAuth, controllers.SendMessage)
	router.GET("/message/:id", middleware.RequireAuth, controllers.GetMessages)
	router.DELETE("/message/:id", middleware.RequireAuth, controllers.DeleteMessage)

	router.Run()
}
