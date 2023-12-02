package postgres

import (
	"github.com/dmytrodemianchuk/go-crud-csv/pkg/db"
	"github.com/jmoiron/sqlx"
)

func Open(cfg db.ConfigDB) (*sqlx.DB, error) {
	return sqlx.Connect("pgx", cfg.String())
}
