package service

import (
	"the_pizza_api/data/request"
	"the_pizza_api/data/response"
	"the_pizza_api/helper"
	"the_pizza_api/model"
	"the_pizza_api/repository"

	"github.com/go-playground/validator/v10"
)

type PizzasService interface {
	Create(pizzas request.CreatePizzasRequest)
	Update(pizzas request.UpdatePizzasRequest)
	Delete(pizzasId uint)
	FindById(pizzasId uint) response.PizzasResponse
	FindAll() []response.PizzasResponse
}

type PizzasServiceImpl struct {
	PizzasRepository repository.PizzasRepository
	Validate         *validator.Validate
}

func NewPizzasServiceImpl(pizzaRepository repository.PizzasRepository, validate *validator.Validate) PizzasService {
	return &PizzasServiceImpl{
		PizzasRepository: pizzaRepository,
		Validate:         validate,
	}
}

// Create implements PizzasService
func (p *PizzasServiceImpl) Create(pizzas request.CreatePizzasRequest) {
	err := p.Validate.Struct(pizzas)
	helper.ErrorPanic(err)
	pizzaModel := model.Pizzas{
		Name:        pizzas.Name,
		Description: pizzas.Description,
	}
	p.PizzasRepository.Save(pizzaModel)
}

// Delete implements PizzasService
func (p *PizzasServiceImpl) Delete(pizzasId uint) {
	p.PizzasRepository.Delete(pizzasId)
}

// FindAll implements PizzasService
func (p *PizzasServiceImpl) FindAll() []response.PizzasResponse {
	result := p.PizzasRepository.FindAll()

	var pizzas []response.PizzasResponse
	for _, value := range result {
		pizza := response.PizzasResponse{
			ID:          value.ID,
			Name:        value.Name,
			Description: value.Description,
		}
		pizzas = append(pizzas, pizza)
	}

	return pizzas
}

// FindById implements PizzasService
func (p *PizzasServiceImpl) FindById(pizzasId uint) response.PizzasResponse {
	pizzaData, err := p.PizzasRepository.FindById(pizzasId)
	helper.ErrorPanic(err)

	pizzaResponse := response.PizzasResponse{
		ID:          pizzaData.ID,
		Name:        pizzaData.Name,
		Description: pizzaData.Description,
	}
	return pizzaResponse
}

// Update implements PizzasService
func (p *PizzasServiceImpl) Update(pizzas request.UpdatePizzasRequest) {
	pizzaData, err := p.PizzasRepository.FindById(pizzas.ID)
	helper.ErrorPanic(err)
	pizzaData.Name = pizzas.Name
	pizzaData.Description = pizzas.Description
	p.PizzasRepository.Update(pizzaData)
}
