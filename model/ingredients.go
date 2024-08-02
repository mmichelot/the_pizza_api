package model

import "gorm.io/gorm"

type Ingredients struct {
	gorm.Model
	Name   string    `gorm:"type:varchar(255)"`
	Pizzas []*Pizzas `gorm:"many2many:pizzas_ingredients;"`
}
