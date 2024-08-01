package service

import (
	"the_pizza_api/data/request"
	"the_pizza_api/data/response"
)

type PizzasService interface {
	Create(pizzas request.CreatePizzasRequest)
	Update(pizzas request.UpdatePizzasRequest)
	Delete(pizzasId int)
	FindById(pizzasId int) response.PizzasResponse
	FindAll() []response.PizzasResponse
}
