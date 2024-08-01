package model

import "gorm.io/gorm"

type Ingredients struct {
	gorm.Model
	Id     int       `gorm:"type:int;primary_key"`
	Name   string    `gorm:"type:varchar(255)"`
	Pizzas []*Pizzas `gorm:"many2many:pizzas_ingredients;"`
}
