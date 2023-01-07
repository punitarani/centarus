package centarus

import (
	"github.com/punitarani/centarus/internal/server"
)

func Run() error {
	mux := server.CreateServer()
	server.RunServer(mux)

	return nil
}
