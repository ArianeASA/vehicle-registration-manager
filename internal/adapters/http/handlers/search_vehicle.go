package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"vehicle-registration-manager/pkg/tracer"
)

// @Title			Search Vehicle
// @Description	Search vehicle By ID
// @Tags			vehicles
// @Accept			json
// @Produce		json
// @Success		200	{object}	responses.Vehicle
// @Failure		404
// @Failure		500
// @Param			id	path	string	true	"Vehicle ID"
// @Router			/vehicles/{id} [get]
func (h *VehicleHandler) HandleSearchVehicleByID(w http.ResponseWriter, r *http.Request) {
	trc := tracer.NewTracer(r)
	vars := mux.Vars(r)
	id := vars["id"]
	if strings.EqualFold("", id) {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	trc.Logger.Infof("Received request id %s", id)

	domain, err := h.searchVehicle.Execute(id)
	if err != nil {
		trc.Logger.Errorf("Failed to list vehicles", nil, err)
		http.Error(w, "Failed to list vehicles", http.StatusInternalServerError)
		return
	}

	if !domain.Exist() {
		trc.Logger.Errorf("Vehicle not found", nil)
		http.Error(w, "Vehicle not found", http.StatusNotFound)
		return
	}

	trc.Logger.Info("Vehicle found")
	vehicles := h.mapDomainToResponseVehicle(domain)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vehicles)
}
