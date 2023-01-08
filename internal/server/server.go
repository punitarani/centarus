package server

import (
	"net/http"
)

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", GetRoot)

	return mux
}
