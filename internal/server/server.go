package server

import (
	"log"
	"net/http"
	"os"
)

var logger = log.New(os.Stdout, "server: ", log.LstdFlags)

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", GetRoot)

	return mux
}
