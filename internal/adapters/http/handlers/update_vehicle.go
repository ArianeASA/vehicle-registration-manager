package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"vehicle-registration-manager/internal/adapters/http/requests"
	"vehicle-registration-manager/pkg/tracer"
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
func (h *VehicleHandler) HandleUpdateVehicle(w http.ResponseWriter, r *http.Request) {
	trc := tracer.NewTracer(r)
	var vehicle requests.Vehicle
	if err := json.NewDecoder(r.Body).Decode(&vehicle); err != nil {
		trc.Logger.Errorf("Failed to decode request body", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	if strings.EqualFold("", id) {
		trc.Logger.Errorf("Bad request", errors.New("id is empty"))
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	domain := h.mapRequestVehicleToDomainVehicle(vehicle)
	domain.ID = id

	trc.Logger.Infof("Received request body %+v", vehicle)
	if err := h.updateVehicle.Execute(domain); err != nil {
		trc.Logger.Errorf("Failed to update vehicle", err)
		http.Error(w, "Failed to update vehicle", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
