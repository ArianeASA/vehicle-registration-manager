package configs

import (
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterHealthCheckRoutes(router *mux.Router) {
	router.HandleFunc("/health", HealthCheck).Methods("GET")
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" && r.Method != "HEAD" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Service is healthy!"))
}
