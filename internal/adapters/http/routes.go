package http

import (
	"github.com/gorilla/mux"
	"net/http"
	"vehicle-registration-manager/internal/adapters/http/handlers"
)

func RegisterRoutes(router *mux.Router, handler *handlers.VehicleHandler) {

	router.HandleFunc("/vehicles", handler.HandleListVehicles).Methods(http.MethodGet)
	router.HandleFunc("/vehicles/{id}", handler.HandleSearchVehicleByID).Methods(http.MethodGet)
	router.HandleFunc("/vehicles/{id}", handler.HandleUpdateVehicle).Methods(http.MethodPut)
	router.HandleFunc("/vehicles/register", handler.HandleRegisterVehicle).Methods(http.MethodPost)

}
