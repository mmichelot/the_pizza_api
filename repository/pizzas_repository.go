package repository

import (
	"errors"
	"the_pizza_api/data/request"
	"the_pizza_api/helper"
	"the_pizza_api/model"

	"gorm.io/gorm"
)

type PizzasRepository interface {
	Save(pizzas model.Pizzas)
	Update(pizzas model.Pizzas)
	Delete(pizzasId uint)
	FindById(pizzasId uint) (pizzas model.Pizzas, err error)
	FindAll() []model.Pizzas
}

type PizzasRepositoryImpl struct {
	Db *gorm.DB
}

func NewPizzasRepositoryImpl(Db *gorm.DB) PizzasRepository {
	return &PizzasRepositoryImpl{Db: Db}
}

// Delete implements PizzasRepository.
func (p *PizzasRepositoryImpl) Delete(pizzasId uint) {
	var pizza model.Pizzas
	result := p.Db.Where("id = ?", pizzasId).Delete(&pizza)
	helper.ErrorPanic(result.Error)
}

// FindAll implements PizzasRepository.
func (p *PizzasRepositoryImpl) FindAll() []model.Pizzas {
	var pizzas []model.Pizzas
	result := p.Db.Find(&pizzas)
	helper.ErrorPanic(result.Error)
	return pizzas
}

// FindById implements PizzasRepository.
func (p *PizzasRepositoryImpl) FindById(pizzasId uint) (pizzas model.Pizzas, err error) {
	var pizza model.Pizzas
	result := p.Db.Find(&pizza, pizzasId)
	if result != nil {
		return pizza, nil
	} else {
		return pizza, errors.New("pizza is not found")
	}
}

// Save implements PizzasRepository.
func (p *PizzasRepositoryImpl) Save(pizza model.Pizzas) {
	result := p.Db.Create(&pizza)
	helper.ErrorPanic(result.Error)
}

// Update implements PizzasRepository.
func (p *PizzasRepositoryImpl) Update(pizza model.Pizzas) {
	var updatePizza = request.UpdatePizzasRequest{
		ID:          pizza.ID,
		Name:        pizza.Name,
		Description: pizza.Description,
	}
	result := p.Db.Model(&pizza).Updates(updatePizza)
	helper.ErrorPanic(result.Error)
}
