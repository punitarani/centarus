package centarus

import (
	"github.com/punitarani/centarus/internal/server"
)

func Run() error {
	mux := server.GetMux()
	server.RunServer(mux)

	return nil
}
