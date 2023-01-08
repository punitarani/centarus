package centarus

import (
	"net/http"

	"github.com/punitarani/centarus/internal/server"
)

func Run() error {
	// Get the mux with the routes registered
	mux := server.GetMux()

	// Listen and serve mux on port 8080
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}

	return nil
}
