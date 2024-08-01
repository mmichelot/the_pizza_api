package controller

import (
	"net/http"
	"strconv"
	"the_pizza_api/data/request"
	"the_pizza_api/data/response"
	"the_pizza_api/helper"
	"the_pizza_api/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type PizzasController struct {
	pizzasService service.PizzasService
}

func NewPizzasController(service service.PizzasService) *PizzasController {
	return &PizzasController{
		pizzasService: service,
	}
}

// CreatePizzas		godoc
// @Summary			Create pizzas
// @Description		Save pizzas data in Db.
// @Param			pizzas body request.CreatePizzasRequest true "Create pizzas"
// @Produce			application/json
// @Pizzas			pizzas
// @Success			200 {object} response.Response{}
// @Router			/pizzas [post]
func (controller *PizzasController) Create(ctx *gin.Context) {
	log.Info().Msg("create pizzas")
	createPizzasRequest := request.CreatePizzasRequest{}
	err := ctx.ShouldBindJSON(&createPizzasRequest)
	helper.ErrorPanic(err)

	controller.pizzasService.Create(createPizzasRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// UpdatePizzas		godoc
// @Summary			Update pizzas
// @Description		Update pizzas data.
// @Param			pizzaId path string true "update pizzas by id"
// @Param			pizzas body request.CreatePizzasRequest true  "Update pizzas"
// @Pizzas			pizzas
// @Produce			application/json
// @Success			200 {object} response.Response{}
// @Router			/pizzas/{pizzaId} [patch]
func (controller *PizzasController) Update(ctx *gin.Context) {
	log.Info().Msg("update pizzas")
	updatePizzasRequest := request.UpdatePizzasRequest{}
	err := ctx.ShouldBindJSON(&updatePizzasRequest)
	helper.ErrorPanic(err)

	pizzaId := ctx.Param("pizzaId")
	id, err := strconv.Atoi(pizzaId)
	helper.ErrorPanic(err)
	updatePizzasRequest.Id = id

	controller.pizzasService.Update(updatePizzasRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// DeletePizzas		godoc
// @Summary			Delete pizzas
// @Description		Remove pizzas data by id.
// @Produce			application/json
// @Pizzas			pizzas
// @Success			200 {object} response.Response{}
// @Router			/pizzas/{pizzaID} [delete]
func (controller *PizzasController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete pizzas")
	pizzaId := ctx.Param("pizzaId")
	id, err := strconv.Atoi(pizzaId)
	helper.ErrorPanic(err)
	controller.pizzasService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindByIdPizzas 		godoc
// @Summary				Get Single pizzas by id.
// @Param				pizzaId path string true "update pizzas by id"
// @Description			Return the tahs whoes pizzaId valu mathes id.
// @Produce				application/json
// @Pizzas				pizzas
// @Success				200 {object} response.Response{}
// @Router				/pizzas/{pizzaId} [get]
func (controller *PizzasController) FindById(ctx *gin.Context) {
	log.Info().Msg("findbyid pizzas")
	pizzaId := ctx.Param("pizzaId")
	id, err := strconv.Atoi(pizzaId)
	helper.ErrorPanic(err)

	pizzaResponse := controller.pizzasService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   pizzaResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAllPizzas 		godoc
// @Summary			Get All pizzas.
// @Description		Return list of pizzas.
// @Pizzas			pizzas
// @Success			200 {obejct} response.Response{}
// @Router			/pizzas [get]
func (controller *PizzasController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll pizzas")
	pizzaResponse := controller.pizzasService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   pizzaResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
