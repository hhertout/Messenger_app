package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	UserSender    uint
	UserRecipient uint
	Message       string
}
