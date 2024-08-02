package main

import (
	"net/http"
	"the_pizza_api/config"
	"the_pizza_api/controller"
	"the_pizza_api/helper"
	"the_pizza_api/middleware"
	"the_pizza_api/model"
	"the_pizza_api/repository"
	"the_pizza_api/router"
	"the_pizza_api/service"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msg("Error loading .env file")
	}
}

func main() {
	log.Info().Msg("Started Server!")

	// Database
	db := config.DatabaseConnection()
	validate := validator.New()

	m := middleware.NewMiddleware(db)

	db.AutoMigrate(&model.Ingredients{}, &model.Pizzas{}, &model.Users{})

	// Repository
	pizzasRepository := repository.NewPizzasRepositoryImpl(db)
	usersRepository := repository.NewUsersRepositoryImpl(db)

	// Service
	pizzasService := service.NewPizzasServiceImpl(pizzasRepository, validate)
	usersService := service.NewUsersServiceImpl(usersRepository, validate)

	// Controller
	pizzasController := controller.NewPizzasController(pizzasService)
	usersController := controller.NewUsersController(usersService)

	routes := router.NewRouter()
	router.PizzasRouter(routes, pizzasController)
	router.UsersRouter(routes, usersController, m)

	server := &http.Server{
		Addr:    ":8000",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
