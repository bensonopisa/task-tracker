package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `gorm:"uniqueIndex;size:30;not null" json:"email"`
	Password  string `gorm:"not null" json:"password"`
	RoleID    uint   `json:"role_id"`
	Tasks     []Task `json:"tasks"`
}
