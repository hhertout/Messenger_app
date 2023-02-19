package datafixtures

import (
	"API_go/go_test/config"
	"API_go/go_test/models"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var userNb = 10

func UsersFixtures() {
	password := "password"

	for i := 0; i < userNb; i++ {
		hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
		if err != nil {
			fmt.Printf("Error during hash => %v", err)
		}
		userMail := fmt.Sprint("user.", i+1, "@gmail.com")
		user := models.User{
			Email:     userMail,
			Password:  string(hash),
			Firstname: "John",
			Lastname:  "Doe",
		}
		result := config.DB.Create(&user)

		if result.Error != nil {
			fmt.Printf("Error during loading fixtures = %v", result.Error)
			return
		}
	}
}
