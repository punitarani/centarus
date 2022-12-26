package sqlstore

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/punitarani/centarus/pkg/setting"
)

type DbSession struct {
	Db *sqlx.DB // Database connection
}

// startSession starts a new database session.
func startSession(cfg *setting.DbCfg) (*DbSession, error) {
	// Establish a new connection to the database.
	db, err := sqlx.Connect("pgx", cfg.DSN)
	if err != nil {
		return nil, err
	}

	// Build the session.
	sess := &DbSession{
		Db: db,
	}

	return sess, nil
}

// closeSession closes the database session.
//
// panic: If the session cannot be closed.
func closeSession(db *DbSession) {
	err := db.Db.Close()
	if err != nil {
		panic(err)
	}
}
