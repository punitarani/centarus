package sqlstore

import (
	"os"
	"testing"

	"github.com/punitarani/centarus/pkg/setting"
)

func TestSQLStore_WithSession(t *testing.T) {
	if os.Getenv("CI") == "true" {
		t.Skip("Skipping test in CI environment")
	}

	// Load config file
	cfg, err := setting.LoadConfigFile("config.toml")
	if err != nil {
		t.Fatalf("Failed to load config file: %v", err)
	}
	dbCfg := cfg.Db["Postgres"]

	// Create a new SQLStore
	ss := &SQLStore{
		Cfg: &dbCfg,
	}

	// Simple TransactionFunc to check if the session is working
	fn := func(sess *DbSession) error {
		// Ping the database
		err := sess.Db.Ping()
		if err != nil {
			return err
		}

		// Get the database time with a query
		var time string
		err = sess.Db.Get(&time, "SELECT NOW()")
		if err != nil {
			return err
		}

		return nil
	}

	// Call the callback with the session
	err = ss.WithSession(fn)
	if err != nil {
		t.Fatal(err)
	}
}
