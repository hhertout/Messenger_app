package migration

import (
	"API_go/go_test/config"
	"API_go/go_test/models"
)

func Migrate() {
	config.DB.AutoMigrate(&models.User{})
}
