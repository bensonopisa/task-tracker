package model

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Name  string  `gorm:"uniqueIndex" json:"name"`
	Users []*User `gorm:"many2many:role_permissions;" json:"users"`
}
