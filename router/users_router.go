package router

import (
	"the_pizza_api/controller"
	"the_pizza_api/middleware"

	"github.com/gin-gonic/gin"
)

func UsersRouter(routes *gin.Engine, usersController *controller.UsersController, m *middleware.Middleware) {
	routes.POST("/signup", usersController.Signup)
	routes.POST("/login", usersController.Login)
	routes.GET("/check", m.Auth, usersController.Check)
}
