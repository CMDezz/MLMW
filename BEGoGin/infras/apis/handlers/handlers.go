package handlers

import (
	"MLMW/BEGoGin/infras/apis/queries"
	"MLMW/BEGoGin/infras/auth"

	"github.com/jmoiron/sqlx"
)

type Handler struct {
	query      queries.Query
	tokenMaker *auth.TokenMaker
}

func NewHandler(store *sqlx.DB, secretCode string, tokenMaker *auth.TokenMaker) Handler {
	return Handler{
		query:      queries.NewQuery(store),
		tokenMaker: tokenMaker,
	}
}
