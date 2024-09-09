package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string       `gorm:"unique" json:"name"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions"`
	Users       []*User      `json:"users"`
}
