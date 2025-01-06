package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"vehicle-registration-manager/internal/adapters/http/requests"
	httpErrors "vehicle-registration-manager/pkg/http_errors"
	"vehicle-registration-manager/pkg/tracer"
)

// @Title			Register Vehicle
// @Description	Register vehicle
// @Summary		Register Vehicle
// @Tags			vehicles
// @Accept			json
// @Produce		json
// @Failure		500 {object} http_errors.ProblemDetails
// @Failure		400 {object} http_errors.ProblemDetails
// @Success		201
// @Param			vehicle	body	requests.Vehicle	true	"Object Vehicle"	example({"brand":"string","model":"string","year":2022,"color":"string","price":4744.32,"license_plate":"string"})
// @Router			/vehicles/register [post]
func (h *vehicleHandler) HandleCreateVehicle(w http.ResponseWriter, r *http.Request) {
	trc := tracer.NewTracer(r)
	var vehicle requests.Vehicle
	if err := json.NewDecoder(r.Body).Decode(&vehicle); err != nil {
		msg := "Failed to decode request body"
		trc.Logger.Error(msg, err)
		httpErrors.WriteProblemDetails(w, httpErrors.BadRequest(msg, r.RequestURI))
		return
	}
	trc.Logger.Infof("Received request body %+v", vehicle)
	domain := h.mapNewRequestVehicleToDomainVehicle(vehicle)
	if !domain.IsValidCreate() {
		msg := "Invalid request body"
		trc.Logger.Error(msg, fmt.Errorf("invalid request body fields [body: %+v]", vehicle))
		httpErrors.WriteProblemDetails(w, httpErrors.BadRequest(msg, r.RequestURI))
		return
	}
	if err := h.createVehicle.Execute(trc, domain); err != nil {
		msg := "Failed to register vehicle"
		trc.Logger.Error(msg, err)
		httpErrors.WriteProblemDetails(w, httpErrors.InternalServerError(msg, r.RequestURI))
		return
	}
	trc.Logger.Info("Vehicle registered")
	w.WriteHeader(http.StatusCreated)
}
