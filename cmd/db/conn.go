package db

import (
	"context"
	"errors"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/punitarani/centarus/pkg/dotenv"
)

// ConnPool is a map of database connection pools.
var ConnPool = make(map[string]*pgxpool.Pool)

// dbs is a map of database names to database urls.
var dbs = make(map[string]string)

// dbEnvNames is a map of database names to environment variable names.
var dbEnvNames = map[string]string{
	"userdata": "DB_USERDATA_URL",
}

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
	// Get the absolute file path of db.env.
	_, file, _, _ := runtime.Caller(0)
	dbEnv := path.Join(filepath.Dir(file), "db.env")

	// Load environment variables.
	if err := dotenv.Load(dbEnv); err != nil {
		return err
	}
	// Get and validate the database urls.
	for name, env := range dbEnvNames {
		url := os.Getenv(env)
		if url == "" {
			return errors.New("missing database url for " + name)
		}
		// Add database url to dbs map.
		dbs[name] = url
	}

	return nil
}
