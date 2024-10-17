package controllers

import (
	"MLMW/BEGoGin/models"
	"MLMW/BEGoGin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller Controller) GetAllPublicTracksController(ctx *gin.Context) {
	res, err := controller.handler.GetAllPublicTrackHandler(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessResponse(res))
}

// TODO: if have time
// GET id user/username from authorization token parsed to payload
// instead of get from uri
func (controller Controller) GetAllTracksByUserIdController(ctx *gin.Context) {
	// var req models.ByIdRequest

	// //validate req
	// if err := ctx.ShouldBindUri(&req); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
	// 	return
	// }
	//-> get payload id from ctx
	Payload := GetPayloadFromCtx(ctx)

	res, err := controller.handler.GetAllTracksByUserIdHandler(ctx, Payload.UserId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessResponse(res))
}

func (controller Controller) CreateTrackController(ctx *gin.Context) {
	var req models.CreateTrackFormRequest

	// Bind form
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	Payload := GetPayloadFromCtx(ctx)

	res, err := controller.handler.CreateTrackHandler(ctx, req, Payload.UserId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessResponse(res))
}

func (controller Controller) UpdateTrackController(ctx *gin.Context) {
	var req models.UpdateTrackFormRequest

	// Bind form
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := controller.handler.UpdateTrackHandler(ctx, req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessResponse(res))
}

func (controller Controller) DeleteTrackByIdController(ctx *gin.Context) {
	var req models.ByIdRequest

	//validate req
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	res, err := controller.handler.DeleteTrackByIdHandler(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessResponse(res))
}

func (controller Controller) GetTrackByIdController(ctx *gin.Context) {
	var req models.ByIdRequest

	//validate req
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	res, err := controller.handler.GetTrackByIdHandler(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessResponse(res))
}
