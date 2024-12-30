package handlers

import (
	"vehicle-registration-manager/internal/adapters/http/mappers"
	"vehicle-registration-manager/internal/adapters/http/requests"
	"vehicle-registration-manager/internal/adapters/http/responses"
	"vehicle-registration-manager/internal/app/usecase/create"
	"vehicle-registration-manager/internal/app/usecase/list"
	"vehicle-registration-manager/internal/app/usecase/search"
	"vehicle-registration-manager/internal/app/usecase/update"
	"vehicle-registration-manager/internal/core/domains"
)

type VehicleHandler struct {
	createVehicle create.CreateVehicle
	updateVehicle update.UpdateVehicle
	listVehicles  list.ListVehicles
	searchVehicle search.SearchVehicle
}

func NewVehicleHandler(
	createVehicle create.CreateVehicle,
	updateVehicle update.UpdateVehicle,
	listVehicles list.ListVehicles,
	searchVehicle search.SearchVehicle,
) *VehicleHandler {
	return &VehicleHandler{
		createVehicle: createVehicle,
		updateVehicle: updateVehicle,
		listVehicles:  listVehicles,
		searchVehicle: searchVehicle,
	}
}

func (h *VehicleHandler) mapDomainToResponseVehicle(domains domains.Vehicle) responses.Vehicle {
	return mappers.DomainVehicleToResponseVehicle(domains)
}

func (h *VehicleHandler) mapRequestVehicleToDomainVehicle(vehicle requests.Vehicle) domains.Vehicle {
	return mappers.RequestVehicleToDomainVehicle(vehicle)
}
func (h *VehicleHandler) mapNewRequestVehicleToDomainVehicle(vehicle requests.Vehicle) domains.Vehicle {
	return mappers.NewRequestVehicleToDomainVehicle(vehicle)
}

func (h *VehicleHandler) mapDomainVehiclesToResponseVehicles(domains []domains.Vehicle) []responses.Vehicle {
	return mappers.DomainVehiclesToResponseVehicles(domains)
}
