package router

import (
	"API_go/go_test/controllers"
	"API_go/go_test/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes() {
	router := gin.New()
	router.SetTrustedProxies([]string{"*"})
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))
	//router.Use(cors.Default())

	// Auth
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)

	// Contacts
	router.GET("/contacts", middleware.RequireAuth, controllers.GetContacts)
	router.GET("/contact/:id", middleware.RequireAuth, controllers.ShowContact)

	router.Run()
}
