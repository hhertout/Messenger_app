package models

import "gorm.io/gorm"

type Invitation struct {
	gorm.Model
	UserSendeur   uint
	UserRecipient uint
	Status        string
}
