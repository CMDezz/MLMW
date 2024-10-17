package controllers

import (
	"MLMW/BEGoGin/models"
	"MLMW/BEGoGin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller Controller) GetAllPublicPlaylists(ctx *gin.Context) {
	res, err := controller.handler.GetAllPublicPlaylistsHandler(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessResponse(res))
}

// TODO: if have time
// GET id user/username from authorization token parsed to payload
// instead of get from uri
func (controller Controller) GetAllPlaylistsByUserIdController(ctx *gin.Context) {
	var req models.ByIdRequest

	//validate req
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	res, err := controller.handler.GetAllPlaylistsByUserIdHandler(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessResponse(res))
}

func (controller Controller) CreatePlaylistController(ctx *gin.Context) {
	var req models.CreatePlaylistFormRequest

	// Bind form
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := controller.handler.CreatePlaylistHandler(ctx, req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessResponse(res))
}

func (controller Controller) UpdatePlaylistController(ctx *gin.Context) {
	var req models.UpdatePlaylistFormRequest

	// Bind form
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := controller.handler.UpdatePlaylistHandler(ctx, req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessResponse(res))
}

func (controller Controller) DeletePlaylistByIdController(ctx *gin.Context) {
	var req models.ByIdRequest

	//validate req
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	res, err := controller.handler.DeletePlaylistByIdHandler(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessResponse(res))
}
