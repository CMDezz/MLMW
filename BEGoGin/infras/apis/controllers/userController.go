package controllers

import (
	"MLMW/BEGoGin/models"
	"MLMW/BEGoGin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller Controller) CreateUserController(ctx *gin.Context) {
	var req models.CreateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	res, err := controller.handler.CreateUserHandler(ctx, req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessResponse(res))
}

func (controller Controller) LoginUserController(ctx *gin.Context) {
	var req models.LoginUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	res, err := controller.handler.LoginUserHandler(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessResponse(res))
}

func (controller Controller) TestUserController(ctx *gin.Context) {
	controller.handler.TestUserHandler(ctx)
	ctx.JSON(http.StatusOK, nil)
}
