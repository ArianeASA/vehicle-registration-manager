package configs

import (
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterHealthCheckRoutes(router *mux.Router) {
	router.HandleFunc("/health", HealthCheck).Methods(http.MethodGet)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Service is healthy!"))
}
