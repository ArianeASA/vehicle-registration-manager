package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"vehicle-registration-manager/internal/adapters/http/requests"
)

// @Title			Update Vehicle
// @Summary		update a task status
// @Description	Update vehicle
// @Tags			vehicles
// @Accept			json
// @Produce		json
// @Success		200
// @Param			id		path	string				false	"Vehicle ID"
// @Param			vehicle	body	requests.Vehicle	true	"Vehicle"	example({"brand":"string","model":"string","year":2022,"color":"string","price":474432})
// @Router			/vehicles/{id} [put]
func (h *VehicleHandler) handleUpdateVehicle(w http.ResponseWriter, r *http.Request) {

	var vehicle requests.Vehicle
	if err := json.NewDecoder(r.Body).Decode(&vehicle); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	if strings.EqualFold("", id) {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	domain := h.mapRequestVehicleToDomainVehicle(vehicle)
	domain.ID = id
	if err := h.updateVehicle.Execute(domain); err != nil {
		http.Error(w, "Failed to update vehicle", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
