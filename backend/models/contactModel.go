package models

import (
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	UserOwnerID   uint
	UserContactID uint
}
