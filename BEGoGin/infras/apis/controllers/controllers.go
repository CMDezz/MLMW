package controllers

import (
	"MLMW/BEGoGin/infras/apis/handlers"
	"MLMW/BEGoGin/infras/auth"

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
