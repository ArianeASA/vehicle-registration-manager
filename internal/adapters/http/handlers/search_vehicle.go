package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	httpErrors "vehicle-registration-manager/pkg/http_errors"
	"vehicle-registration-manager/pkg/tracer"
)

// @Title			Search Vehicle
// @Description	Search vehicle By ID
// @Summary		Search Vehicle by ID
// @Tags			vehicles
// @Accept			json
// @Produce		json
// @Success		200	{object}	responses.Vehicle
// @Failure		404	{object}	http_errors.ProblemDetails
// @Failure		500	{object}	http_errors.ProblemDetails
// @Failure		400	{object}	http_errors.ProblemDetails
// @Param			id	path		string	true	"Vehicle ID"
// @Router			/vehicles/{id} [get]
func (h *vehicleHandler) HandleSearchVehicleByID(w http.ResponseWriter, r *http.Request) {
	trc := tracer.NewTracer(r)
	vars := mux.Vars(r)
	id := vars["id"]
	if strings.EqualFold("", id) {
		msg := "Invalid id"
		httpErrors.WriteProblemDetails(w, httpErrors.BadRequest(msg, r.RequestURI))
		return
	}

	trc.Logger.Infof("Received request id %s", id)

	domain, err := h.searchVehicle.Execute(trc, id)
	if err != nil {
		msg := "Failed to search vehicles"
		trc.Logger.Error(msg, err)
		httpErrors.WriteProblemDetails(w, httpErrors.InternalServerError(msg, r.RequestURI))
		return
	}

	if !domain.Exist() {
		msg := "Vehicle not found"
		trc.Logger.Errorf(msg, nil)
		httpErrors.WriteProblemDetails(w, httpErrors.NotFound(msg, r.RequestURI))
		return
	}

	trc.Logger.Info("Vehicle found")
	vehicles := h.mapDomainToResponseVehicle(domain)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vehicles)
}
