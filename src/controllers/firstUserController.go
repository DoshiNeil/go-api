package controllers

import (
	"go-api/src/models"
	"go-api/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type firstUserController struct {
	service services.UserService
}

func NewFirstUserController(service services.UserService) Controller {
	return &firstUserController{
		service: service,
	}
}

// Handler implements Controller.
func (f *firstUserController) Handler() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		body := &models.User{}
		err := ctx.BindJSON(body)
		if err != nil {
			ctx.JSON(404, err)
		}

	}
}

// Method implements Controller.
func (f *firstUserController) Method() string {
	return http.MethodGet
}

// Route implements Controller.
func (f *firstUserController) Route() string {
	return "/firstUser"
}
