package server

import (
	"net/http"
)

func CreateServer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", GetRoot)

	return mux
}

func RunServer(mux *http.ServeMux) {
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
