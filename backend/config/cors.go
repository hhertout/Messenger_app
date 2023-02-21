package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSConfig(router *gin.Engine) {
	router.SetTrustedProxies([]string{"*"})
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowCredentials = true
	router.Use(cors.New(config))
}
