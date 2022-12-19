package user

import (
	"context"
	"errors"
	"github.com/punitarani/centarus/pkg/db"

	"github.com/punitarani/centarus/pkg/random"
)

const UIDLength = 30

// GenUID generates a unique identifier for a user.
// UID is a 30 character string of alphanumeric characters.
func GenUID() (string, error) {
	ctx := context.Background()
	userdataDB := db.ConnPool["userdata"]
	if userdataDB == nil {
		return "", errors.New("userdata database not connected")
	}

	// Generate a UID that is not already in the database.
	for {
		uid, err := random.String(UIDLength, random.AlphaNumeric)
		if err != nil {
			return "", err
		}

		// Query the database to see if the UID is already in use.
		// If UID already exists, generate a new one.
		var count int
		err = userdataDB.QueryRow(ctx, "SELECT COUNT(*) FROM user_info WHERE user_id = $1", uid).Scan(&count)
		if err != nil {
			return "", err
		}

		// If UID is unique, return it.
		if count == 0 {
			return uid, nil
		}
	}
}
