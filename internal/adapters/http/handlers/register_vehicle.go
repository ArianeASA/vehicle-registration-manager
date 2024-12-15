package handlers

import (
	"encoding/json"
	"net/http"
	"vehicle-registration-manager/internal/adapters/http/requests"
)

// @Title			Register Vehicle
// @Description	Register vehicle
// @Tags			vehicles
// @Accept			json
// @Produce		json
// @Success		201
// @Param			vehicle	body	requests.Vehicle	true	"Object Vehicle"	example({"brand":"string","model":"string","year":2022,"color":"string","price":4744.32})
// @Router			/vehicles/register [post]
func (h *VehicleHandler) handleRegisterVehicle(w http.ResponseWriter, r *http.Request) {

	var vehicle requests.Vehicle
	if err := json.NewDecoder(r.Body).Decode(&vehicle); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	domain := h.mapNewRequestVehicleToDomainVehicle(vehicle)
	if err := h.registerVehicle.Execute(domain); err != nil {
		http.Error(w, "Failed to register vehicle", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
