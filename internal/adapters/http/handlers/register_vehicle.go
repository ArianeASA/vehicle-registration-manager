package handlers

import (
	"encoding/json"
	"net/http"
	"vehicle-registration-manager/internal/adapters/http/requests"
	"vehicle-registration-manager/internal/core/domain"
)

func (h *VehicleHandler) handleRegisterVehicle(w http.ResponseWriter, r *http.Request) {

	var vehicle requests.Vehicle
	if err := json.NewDecoder(r.Body).Decode(&vehicle); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	vehicleDomain := domain.Vehicle{
		Brand: vehicle.Brand,
		Model: vehicle.Model,
		Year:  vehicle.Year,
		Color: vehicle.Color,
		Price: vehicle.Price,
	}
	if err := h.registerVehicle.Execute(vehicleDomain); err != nil {
		http.Error(w, "Failed to register vehicle", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
