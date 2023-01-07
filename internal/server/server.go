package server

import (
	"net/http"
)

func CreateServer() *http.ServeMux {
	mux := http.NewServeMux()

	return mux
}

func RunServer(mux *http.ServeMux) {
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
