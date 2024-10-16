package controllers

import (
	"MLMW/BEGoGin/infras/apis/handlers"

	"github.com/jmoiron/sqlx"
)

type Controller struct {
	handler handlers.Handler
}

func NewController(store *sqlx.DB) Controller {
	return Controller{
		handler: handlers.NewHandler(store),
	}
}
