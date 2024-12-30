package handlers_test

import (
	"github.com/gorilla/mux"
	"vehicle-registration-manager/internal/adapters/http/handlers"
	"vehicle-registration-manager/internal/app/usecase/create"
	"vehicle-registration-manager/internal/app/usecase/list"
	"vehicle-registration-manager/internal/app/usecase/search"
	"vehicle-registration-manager/internal/app/usecase/update"
)

var (
	handler       *handlers.VehicleHandler
	router        *mux.Router
	listVehicles  *list.MockListVehicles
	createVehicle *create.MockCreateVehicle
	updateVehicle *update.MockUpdateVehicle
	searchVehicle *search.MockSearchVehicle
)

func setup() {
	listVehicles = new(list.MockListVehicles)
	createVehicle = new(create.MockCreateVehicle)
	updateVehicle = new(update.MockUpdateVehicle)
	searchVehicle = new(search.MockSearchVehicle)
	handler = handlers.NewVehicleHandler(createVehicle, updateVehicle, listVehicles, searchVehicle)
	router = mux.NewRouter()
}
