package connections

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitDBConnection(dbDriver string, dbSource string) (*sqlx.DB, error) {
	// make connection to Postgresql by SQLX
	sqlxDB, err := sqlx.Open(dbDriver, dbSource)
	if err != nil {
		return nil, err
	}

	err = sqlxDB.Ping()
	if err != nil {
		return nil, err
	}

	return sqlxDB, nil
}
