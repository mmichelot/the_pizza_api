package controller

import (
	"net/http"

	"the_pizza_api/data/request"
	"the_pizza_api/data/response"
	"the_pizza_api/helper"
	"the_pizza_api/model"
	"the_pizza_api/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type UsersController struct {
	usersService service.UsersService
}

func NewUsersController(service service.UsersService) *UsersController {
	return &UsersController{
		usersService: service,
	}
}

func (controller *UsersController) Signup(ctx *gin.Context) {
	log.Info().Msg("Signup users")

	createUsersRequest := request.CreateUsersRequest{}

	err := ctx.ShouldBindJSON(&createUsersRequest)
	helper.ErrorPanic(err)

	controller.usersService.Create(createUsersRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UsersController) Login(ctx *gin.Context) {
	log.Info().Msg("Signup users")

	createUsersRequest := request.CreateUsersRequest{}

	err := ctx.ShouldBindJSON(&createUsersRequest)
	helper.ErrorPanic(err)

	tokenString := controller.usersService.Login(createUsersRequest)

	// set to cookie
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600, "", "", true, true)

	// send success response
	ctx.JSON(http.StatusOK, gin.H{})
}

func (controller *UsersController) Check(c *gin.Context) {
	user, exist := c.Get("user")
	if !exist {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not exists",
		})
		return
	}

	email := user.(model.Users).Email
	c.JSON(http.StatusOK, gin.H{
		"message": map[string]string{"email": email},
	})
}
