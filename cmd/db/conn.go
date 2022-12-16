package db

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/punitarani/centarus/pkg/dotenv"
)

// ConnPool is a map of database connection pools.
var ConnPool = make(map[string]*pgxpool.Pool)

// dbEnvNames is a map of database names to environment variable names.
var dbEnvNames = map[string]string{
	"userdata": "DB_USERDATA_URL",
}

// dbs is a map of database names to database urls.
var dbs = make(map[string]string)

// CreateConnections creates connections to all databases.
func CreateConnections() error {
	if err := loadEnvVars(); err != nil {
		return err
	}

	// Create connection pool for each database.
	for name, url := range dbs {
		pool, err := pgxpool.New(context.Background(), url)
		if err != nil {
			return err
		}
		ConnPool[name] = pool
	}

	return nil
}

// CloseConnections closes connections to all databases.
func CloseConnections() {
	for name, pool := range ConnPool {
		if pool != nil {
			pool.Close()
		}
		ConnPool[name] = nil
	}
}

// loadEnvVars loads environment variables from the db.env file.
func loadEnvVars() error {
	// Load environment variables.
	if err := dotenv.Load("db.env"); err != nil {
		return err
	}

	// Check that all database urls are set.
	for name, url := range dbEnvNames {
		if url == "" {
			return errors.New("missing database url for " + name)
		}
	}

	return nil
}
