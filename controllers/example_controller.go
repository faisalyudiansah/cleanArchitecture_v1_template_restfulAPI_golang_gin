package controllers

import (
	"server/helpers/ginutils"
	"server/services"

	"github.com/gin-gonic/gin"
)

type ExampleController struct {
	ExampleService services.ExampleServiceInterface
}

func NewExampleController(
	ExampleService services.ExampleServiceInterface,
) *ExampleController {
	return &ExampleController{
		ExampleService: ExampleService,
	}
}

func (sc *ExampleController) Get(ctx *gin.Context) {
	err := sc.ExampleService.GetService(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	ginutils.ResponseOKPlain(ctx)
}
