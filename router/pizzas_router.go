package router

import (
	"the_pizza_api/controller"

	"github.com/gin-gonic/gin"
)

func PizzasRouter(routes *gin.Engine, pizzasController *controller.PizzasController) {
	pizzasRouter := routes.Group("/pizzas")
	pizzasRouter.GET("", pizzasController.FindAll)
	pizzasRouter.GET("/:pizzaId", pizzasController.FindById)
	pizzasRouter.POST("", pizzasController.Create)
	pizzasRouter.PATCH("/:pizzaId", pizzasController.Update)
	pizzasRouter.DELETE("/:pizzaId", pizzasController.Delete)
}
