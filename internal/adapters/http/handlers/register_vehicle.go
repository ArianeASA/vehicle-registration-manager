package handlers

import (
	"encoding/json"
	"net/http"
	"vehicle-registration-manager/internal/adapters/http/requests"
	"vehicle-registration-manager/pkg/tracer"
)

// @Title			Register Vehicle
// @Description	Register vehicle
// @Tags			vehicles
// @Accept			json
// @Produce		json
// @Success		201
// @Param			vehicle	body	requests.Vehicle	true	"Object Vehicle"	example({"brand":"string","model":"string","year":2022,"color":"string","price":4744.32})
// @Router			/vehicles/register [post]
func (h *VehicleHandler) HandleRegisterVehicle(w http.ResponseWriter, r *http.Request) {
	trc := tracer.NewTracer(r)
	var vehicle requests.Vehicle
	if err := json.NewDecoder(r.Body).Decode(&vehicle); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		trc.Logger.Errorf("Failed to decode request body", nil, err)
		return
	}
	trc.Logger.Infof("Received request body %+v", vehicle)
	domain := h.mapNewRequestVehicleToDomainVehicle(vehicle)
	if err := h.registerVehicle.Execute(trc, domain); err != nil {
		trc.Logger.Errorf("Failed to register vehicle", nil, err)
		http.Error(w, "Failed to register vehicle", http.StatusInternalServerError)
		return
	}
	trc.Logger.Info("Vehicle registered")
	w.WriteHeader(http.StatusCreated)
}
