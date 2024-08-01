package model

import "gorm.io/gorm"

type Pizzas struct {
	gorm.Model
	Id          int            `gorm:"type:int;primary_key"`
	Name        string         `gorm:"type:varchar(255)"`
	Description string         `gorm:"type:varchar(255)"`
	Ingredients []*Ingredients `gorm:"many2many:pizzas_ingredients;"`
}
