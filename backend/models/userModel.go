package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email            string `gorm:"unique"`
	Password         string
	Firstname        string
	Lastname         string
	Owner            Contact    `gorm:"foreignKey:UserOwnerID"`
	Contact          Contact    `gorm:"foreignKey:UserContactID"`
	Sender           Invitation `gorm:"foreignKey:UserSendeur"`
	Recipient        Invitation `gorm:"foreignKey:UserRecipient"`
	MessageSender    Message    `gorm:"foreignKey:UserSender"`
	MessageRecipient Message    `gorm:"foreignKey:UserRecipient"`
}
