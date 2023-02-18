package datafixtures

import (
	"API_go/go_test/config"
	"API_go/go_test/models"
	"fmt"
	"math/rand"
)

var contactNb = 50

func ContactFixtures() {

	for i := 0; i < contactNb; i++ {
		fmt.Printf("%d", i)
		ownerid := rand.Intn(userNb)
		if ownerid == 0 {
			ownerid = 1
		}
		contactid := rand.Intn(userNb)
		if contactid == 0 {
			contactid = 1
		}

		contact := models.Contact{
			UserOwnerID:   uint(ownerid),
			UserContactID: uint(contactid),
		}
		result := config.DB.Create(&contact)
		if result.Error != nil {
			fmt.Printf("cannot save contact (error: %v)", result.Error)
		}
	}
}
