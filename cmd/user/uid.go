package user

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/punitarani/centarus/pkg/random"
)

const UIDLength = 30

// GenUID generates a unique identifier for a user.
// UID is a 30 character string of alphanumeric characters.
func GenUID() (string, error) {
	// Connect to the database.
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, os.Getenv("DB_USERDATA_URL"))
	if err != nil {
		return "", fmt.Errorf("unable to connect to database: %w", err)
	}
	defer func() {
		if err := conn.Close(ctx); err != nil {
			log.Printf("unable to close database connection: %v", err)
		}
	}()

	// Generate a UID that is not already in the database.
	for {
		uid, err := random.String(UIDLength, random.AlphaNumeric)
		if err != nil {
			return "", err
		}

		// Query the database to see if the UID is already in use.
		// If UID already exists, generate a new one.
		var count int
		err = conn.QueryRow(ctx, "SELECT COUNT(*) FROM user_info WHERE user_id = $1", uid).Scan(&count)
		if err != nil {
			return "", err
		}

		// If UID is unique, return it.
		if count == 0 {
			return uid, nil
		}
	}
}
