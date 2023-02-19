package migration

import (
	"API_go/go_test/config"
	"API_go/go_test/models"
)

func Migrate() {
	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Contact{})
	config.DB.AutoMigrate(&models.Invitation{})
	config.DB.AutoMigrate(&models.Message{})
}
