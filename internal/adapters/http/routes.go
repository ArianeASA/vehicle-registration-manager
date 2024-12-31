package http

import (
	"github.com/gorilla/mux"
	"net/http"
	"vehicle-registration-manager/internal/core/ports/in"
)

func RegisterRoutes(router *mux.Router, handler in.VehicleHandler) {
	router.StrictSlash(true)
	router.HandleFunc("/vehicles", handler.HandleListVehicles).Methods(http.MethodGet).Name("ListVehicles")
	router.HandleFunc("/vehicles/{id}", handler.HandleSearchVehicleByID).Methods(http.MethodGet).Name("SearchVehicleByID")
	router.HandleFunc("/vehicles/{id}", handler.HandleUpdateVehicle).Methods(http.MethodPut).Name("UpdateVehicle")
	router.HandleFunc("/vehicles/register", handler.HandleCreateVehicle).Methods(http.MethodPost).Name("CreateVehicle")
}
