package in

import (
	"net/http"
)

type VehicleHandler interface {
	HandleCreateVehicle(w http.ResponseWriter, r *http.Request)
	HandleUpdateVehicle(w http.ResponseWriter, r *http.Request)
	HandleListVehicles(w http.ResponseWriter, r *http.Request)
	HandleSearchVehicleByID(w http.ResponseWriter, r *http.Request)
}
