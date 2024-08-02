package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
