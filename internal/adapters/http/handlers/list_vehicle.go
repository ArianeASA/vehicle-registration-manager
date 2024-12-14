package handlers

import (
	"encoding/json"
	"net/http"
)

//	@Title			List Vehicles
//	@Description	List vehicles
//	@Tags			vehicles
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]domain.Vehicle
//	@Router			/vehicles [get]
func (h *VehicleHandler) handleListVehicles(w http.ResponseWriter, r *http.Request) {
	vehicles, err := h.listVehicles.Execute()
	if err != nil {
		http.Error(w, "Failed to list vehicles", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vehicles)
}
