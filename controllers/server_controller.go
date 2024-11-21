package controllers

import (
	"server/services"

	"github.com/gin-gonic/gin"
)

type ServerController struct {
	ServerService services.ServerServiceInterface
}

func NewServerController(
	ServerService services.ServerServiceInterface,
) *ServerController {
	return &ServerController{
		ServerService: ServerService,
	}
}

func (sc *ServerController) Get(ctx *gin.Context) {
	err := sc.ServerService.GetService(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
}
