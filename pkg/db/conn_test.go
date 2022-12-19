package db

import (
	"context"
	"testing"
)

func TestCreateCloseConnections(t *testing.T) {
	if err := CreateConnections(); err != nil {
		t.Errorf("unable to create connections: %v", err)
	}
	// This ensures that the connections are closed if the function returns early.
	defer CloseConnections()

	// Check that ConnPool is not empty.
	if len(ConnPool) == 0 {
		t.Errorf("ConnPool is empty")
	}

	// Check that all connections were created.
	for name, pool := range ConnPool {
		if pool == nil {
			t.Errorf("connection pool for %v is nil", name)
		}

		// Make a test query to ensure that the connection is valid.
		for _, pool := range ConnPool {
			if _, err := pool.Exec(context.Background(), "SELECT 1"); err != nil {
				t.Errorf("unable to query database: %v", err)
			}
		}
	}

	// Close connections.
	CloseConnections()

	// Check that all connections were closed.
	for name, pool := range ConnPool {
		if pool != nil {
			t.Errorf("connection pool for %v is not nil", name)
		}
		if ConnPool[name] != nil {
			t.Errorf("ConnPool{%v} is not nil", name)
		}
	}
}
