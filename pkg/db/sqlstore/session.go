package sqlstore

import "github.com/jmoiron/sqlx"

type DbSession struct {
	Db *sqlx.DB // Database connection
}
