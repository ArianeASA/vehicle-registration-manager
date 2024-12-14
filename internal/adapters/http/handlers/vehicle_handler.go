package handlers

import (
	"github.com/gorilla/mux"
	"vehicle-registration-manager/internal/app/usecase"
)

type VehicleHandler struct {
	registerVehicle *usecase.RegisterVehicle
	updateVehicle   *usecase.UpdateVehicle
	listVehicles    *usecase.ListVehicles
}

func NewVehicleHandler(
	registerVehicle *usecase.RegisterVehicle,
	updateVehicle *usecase.UpdateVehicle,
	listVehicles *usecase.ListVehicles,
) *VehicleHandler {
	return &VehicleHandler{
		registerVehicle: registerVehicle,
		updateVehicle:   updateVehicle,
		listVehicles:    listVehicles,
	}
}

func (h *VehicleHandler) RegisterRoutes(router *mux.Router) {

	router.HandleFunc("/vehicles", h.handleListVehicles).Methods("GET")
	router.HandleFunc("/vehicles/register", h.handleRegisterVehicle).Methods("POST")
	router.HandleFunc("/vehicles/update/{id}", h.handleUpdateVehicle).Methods("PUT")
}
