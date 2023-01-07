package server

import "net/http"

func GetRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte("Centarus"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}
}
