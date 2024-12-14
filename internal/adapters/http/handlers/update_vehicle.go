package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"vehicle-registration-manager/internal/adapters/http/requests"
	"vehicle-registration-manager/internal/core/domain"
)

func (h *VehicleHandler) handleUpdateVehicle(w http.ResponseWriter, r *http.Request) {

	var vehicle requests.Vehicle
	if err := json.NewDecoder(r.Body).Decode(&vehicle); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	id := r.PathValue("id")
	if strings.EqualFold("", id) {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	vehicleDomain := domain.Vehicle{
		ID:    id,
		Brand: vehicle.Brand,
		Model: vehicle.Model,
		Year:  vehicle.Year,
		Color: vehicle.Color,
		Price: vehicle.Price,
	}

	if err := h.updateVehicle.Execute(vehicleDomain); err != nil {
		http.Error(w, "Failed to update vehicle", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
