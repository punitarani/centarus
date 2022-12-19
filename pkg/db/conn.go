package db

import (
	"context"
	"errors"
	"os"
	"os/exec"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/punitarani/centarus/pkg/dotenv"
)

// dbEnvNames is a map of database names to environment variable names.
var dbEnvNames = map[string]string{
	"userdata": "DB_USERDATA_URL",
}

// ConnPool is a map of database connection pools.
var ConnPool = make(map[string]*pgxpool.Pool)

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

// loadEnvVars reads the database urls from the environment variables and stores them in the dbs map.
// If .secrets.env is present in project root, it is loaded into the environment.
func loadEnvVars() error {
	// Get the path to the .secrets.env file.
	dbDir, err := exec.Command("go", "list", "-f", "{{.Root}}").Output()
	if err != nil {
		return err
	}
	dbEnv := string(dbDir[:len(dbDir)-1]) + "/.secrets.env"

	// Load environment variables.
	if err := dotenv.Load(dbEnv, false); err != nil {
		return err
	}
	// Get and validate the database urls.
	for name, env := range dbEnvNames {
		url := os.Getenv(env)
		if url == "" {
			return errors.New("missing environment variable: " + env)
		}
		// Add database url to dbs map.
		dbs[name] = url
	}

	return nil
}
