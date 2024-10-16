package queries

import "github.com/jmoiron/sqlx"

type Query struct {
	store *sqlx.DB
}

func NewQuery(store *sqlx.DB) Query {
	return Query{
		store: store,
	}
}
