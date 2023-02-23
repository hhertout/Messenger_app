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
	router.POST("api/signup", controllers.Signup)
	router.POST("api/login", controllers.Login)
	router.POST("api/logout", middleware.RequireAuth, controllers.Logout)
	router.GET("api/validate", middleware.RequireAuth, controllers.Validate)

	// User
	router.GET("api/user/:id", middleware.RequireAuth, controllers.GetUser)

	// Contacts
	router.GET("api/contacts", middleware.RequireAuth, controllers.GetContacts)

	// Invitations
	router.GET("api/invite/pending", middleware.RequireAuth, controllers.GetInvitation)
	router.GET("api/invite/accepted", middleware.RequireAuth, controllers.GetInvitation)
	router.POST("api/invite", middleware.RequireAuth, controllers.Invite)
	router.DELETE("api/invite/:id", middleware.RequireAuth, controllers.DeleteInvitation)
	router.PUT("api/invite/:id", middleware.RequireAuth, controllers.UpdateInvitation)

	// Message
	router.POST("api/message", middleware.RequireAuth, controllers.SendMessage)
	router.GET("api/message/:id", middleware.RequireAuth, controllers.GetMessages)
	router.DELETE("api/message/:id", middleware.RequireAuth, controllers.DeleteMessage)

	router.Run()
}
