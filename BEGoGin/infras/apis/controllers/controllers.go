package controllers

import (
	"MLMW/BEGoGin/infras/apis/handlers"
	"MLMW/BEGoGin/infras/auth"
	"MLMW/BEGoGin/utils"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Controller struct {
	handler handlers.Handler
}

func NewController(store *sqlx.DB, secretCode string, tokenMaker *auth.TokenMaker) Controller {
	return Controller{
		handler: handlers.NewHandler(store, secretCode, tokenMaker),
	}
}

func GetPayloadFromCtx(ctx *gin.Context) auth.Payload {
	pl := ctx.Value(utils.AUTHORIZATION_PAYLOAD_KEY)
	payload, ok := pl.(*auth.Payload)
	if !ok {
		return auth.Payload{}
	}
	return *payload // Return the payload directly (no need for &payload)
}
