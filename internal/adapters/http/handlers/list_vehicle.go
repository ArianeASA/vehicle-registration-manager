package handlers

import (
	"encoding/json"
	"net/http"
	"vehicle-registration-manager/pkg/tracer"
)

// @Title			List Vehicles
// @Description	List vehicles
// @Tags			vehicles
// @Accept			json
// @Produce		json
// @Success		200	{object}	[]responses.Vehicle
// @Router			/vehicles [get]
func (h *VehicleHandler) HandleListVehicles(w http.ResponseWriter, r *http.Request) {
	trc := tracer.NewTracer(r)
	domains, err := h.listVehicles.Execute(trc)
	if err != nil {
		trc.Logger.Errorf("Failed to list vehicles", nil, err)
		http.Error(w, "Failed to list vehicles", http.StatusInternalServerError)
		return
	}

	trc.Logger.Info("List vehicles")
	vehicles := h.mapDomainVehiclesToResponseVehicles(domains)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vehicles)
}
