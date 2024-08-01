package repository

import "the_pizza_api/model"

type PizzasRepository interface {
	Save(pizzas model.Pizzas)
	Update(pizzas model.Pizzas)
	Delete(pizzasId int)
	FindById(pizzasId int) (pizzas model.Pizzas, err error)
	FindAll() []model.Pizzas
}
