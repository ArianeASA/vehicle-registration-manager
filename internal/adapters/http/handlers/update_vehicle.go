package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"vehicle-registration-manager/internal/adapters/http/requests"
	"vehicle-registration-manager/internal/core/domains"
	httpErrors "vehicle-registration-manager/pkg/http_errors"
	"vehicle-registration-manager/pkg/tracer"
)

// @Title			Update Vehicle
// @Summary		update a task status
// @Description	Update vehicle
// @Tags			vehicles
// @Accept			json
// @Produce		json
// @Success		200
// @Failure		404		{object}	http_errors.ProblemDetails
// @Failure		500		{object}	http_errors.ProblemDetails
// @Failure		400		{object}	http_errors.ProblemDetails
// @Param			id		path		string				false	"Vehicle ID"
// @Param			vehicle	body		requests.Vehicle	true	"Vehicle"	example({"brand":"string","model":"string","year":2022,"color":"string","price":474432, "license_plate":"string", "status":"FOR_SALE"})
// @Router			/vehicles/{id} [put]
func (h *vehicleHandler) HandleUpdateVehicle(w http.ResponseWriter, r *http.Request) {
	trc := tracer.NewTracer(r)
	var vehicle requests.Vehicle
	if err := json.NewDecoder(r.Body).Decode(&vehicle); err != nil {
		msg := "Failed to decode request body"
		trc.Logger.Errorf(msg, err)
		httpErrors.WriteProblemDetails(w, httpErrors.BadRequest(msg, r.RequestURI))
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	if strings.EqualFold("", id) {
		msg := "Invalid id"
		trc.Logger.Errorf(msg, errors.New("id is empty"))
		httpErrors.WriteProblemDetails(w, httpErrors.BadRequest(msg, r.RequestURI))
		return
	}

	domain := h.mapRequestVehicleToDomainVehicle(vehicle)
	domain.ID = id

	trc.Logger.Infof("Received request body %+v", vehicle)
	if err := h.updateVehicle.Execute(trc, domain); err != nil {
		handleUpdateError(w, r, trc, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleUpdateError(w http.ResponseWriter, r *http.Request, trc *tracer.Tracer, err error) {
	if errors.Is(err, domains.ErrVehicleNotFound) {
		msg := "Vehicle not found"
		trc.Logger.Errorf(msg, err)
		httpErrors.WriteProblemDetails(w, httpErrors.NotFound(msg, r.RequestURI))
		return
	}

	msg := "Failed to update vehicle"
	trc.Logger.Errorf(msg, err)
	httpErrors.WriteProblemDetails(w, httpErrors.InternalServerError(msg, r.RequestURI))
	return
}
