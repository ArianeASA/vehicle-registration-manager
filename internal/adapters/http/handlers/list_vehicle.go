package handlers

import (
	"encoding/json"
	"net/http"
	httpErrors "vehicle-registration-manager/pkg/http_errors"
	"vehicle-registration-manager/pkg/tracer"
)

// @Title			List Vehicles
// @Description	List vehicles
// @Tags			vehicles
// @Accept			json
// @Produce		json
// @Failure		500 {object} http_errors.ProblemDetails
// @Success		200	{object}	[]responses.Vehicle
// @Router			/vehicles [get]
func (h *vehicleHandler) HandleListVehicles(w http.ResponseWriter, r *http.Request) {
	trc := tracer.NewTracer(r)
	domains, err := h.listVehicles.Execute(trc)
	if err != nil {
		msg := "Failed to list vehicles"
		trc.Logger.Error(msg, err)
		httpErrors.WriteProblemDetails(w, httpErrors.InternalServerError(msg, r.RequestURI))
		return
	}

	vehicles := h.mapDomainVehiclesToResponseVehicles(domains)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vehicles)
}
