package user

import (
	"testing"

	"github.com/punitarani/centarus/cmd/db"
)

func TestGenUID(t *testing.T) {
	// Setup
	if err := db.CreateConnections(); err != nil {
		t.Errorf("unable to create connections: %v", err)
	}
	defer db.CloseConnections()

	var uids []string

	// Generate 2**4 unique UIDs.
	for i := 0; i < 1<<4; i++ {
		uid, err := GenUID()
		if err != nil {
			t.Errorf("unable to generate UID: %v", err)
		}

		// Check that the UID is unique.
		for _, u := range uids {
			if uid == u {
				t.Errorf("duplicate UID generated: %v", uid)
			}
		}
		uids = append(uids, uid)
	}
}
