package server

import "net/http"

func LogRequest(r *http.Request) {
	logger.Println("server: ", r.Method, r.URL.Path, r.URL.Query(), r.Header, r.Body)
}
