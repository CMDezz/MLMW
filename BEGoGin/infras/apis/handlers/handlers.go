package handlers

import (
	"MLMW/BEGoGin/infras/apis/queries"

	"github.com/jmoiron/sqlx"
)

type Handler struct {
	query queries.Query
}

func NewHandler(store *sqlx.DB) Handler {

	return Handler{
		query: queries.NewQuery(store),
	}
}
