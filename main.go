package main

import (
	"net/http"
	"the_pizza_api/config"
	"the_pizza_api/controller"
	"the_pizza_api/helper"
	"the_pizza_api/model"
	"the_pizza_api/repository"
	"the_pizza_api/router"
	"the_pizza_api/service"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Started Server!")
	// Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.AutoMigrate(&model.Ingredients{}, &model.Pizzas{})

	// Repository
	pizzasRepository := repository.NewPizzasRepositoryImpl(db)

	// Service
	pizzasService := service.NewPizzasServiceImpl(pizzasRepository, validate)

	// Controller
	pizzasController := controller.NewPizzasController(pizzasService)

	routes := router.NewRouter()
	router.PizzasRouter(routes, pizzasController)

	server := &http.Server{
		Addr:    ":8000",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
