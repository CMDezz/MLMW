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

func (controller Controller) UpsertTracksPlaylistsController(ctx *gin.Context) {
	var req models.UpsertTracksPlaylistsModel
	//validate req
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := controller.handler.UpsertTracksPlayListsHandler(ctx, req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessResponse(res))
}

func (controller Controller) GetDataTrackPlaylistByIdController(ctx *gin.Context) {
	var req models.ByIdRequest
	//validate req
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := controller.handler.GetDataTrackPlaylistByIdHandler(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessResponse(res))
}

func (controller Controller) GetFullPlaylistDetailController(ctx *gin.Context) {
	var req models.ByIdRequest
	//validate req
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := controller.handler.GetFullPlaylistDetailHandler(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessResponse(res))

}
