package controllers

import (
	"MLMW/BEGoGin/models"
	"MLMW/BEGoGin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Common ControllerHanlder is used for common apis like:
// searching (track+artis)

func (controller Controller) SearchingController(ctx *gin.Context) {
	var req models.SearchingRequest

	//validate req
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := controller.handler.SearchingHandler(ctx, req.Keyword)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessResponse(res))
}
