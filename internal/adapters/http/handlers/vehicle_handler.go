package handlers

import (
	"vehicle-registration-manager/internal/adapters/http/mappers"
	"vehicle-registration-manager/internal/adapters/http/requests"
	"vehicle-registration-manager/internal/adapters/http/responses"
	"vehicle-registration-manager/internal/app/usecase"
	"vehicle-registration-manager/internal/core/domain"
)

type VehicleHandler struct {
	registerVehicle *usecase.RegisterVehicle
	updateVehicle   *usecase.UpdateVehicle
	listVehicles    *usecase.ListVehicles
	searchVehicle   *usecase.SearchVehicle
}

func NewVehicleHandler(
	registerVehicle *usecase.RegisterVehicle,
	updateVehicle *usecase.UpdateVehicle,
	listVehicles *usecase.ListVehicles,
	searchVehicle *usecase.SearchVehicle,
) *VehicleHandler {
	return &VehicleHandler{
		registerVehicle: registerVehicle,
		updateVehicle:   updateVehicle,
		listVehicles:    listVehicles,
		searchVehicle:   searchVehicle,
	}
}

func (h *VehicleHandler) mapDomainToResponseVehicle(domains domain.Vehicle) responses.Vehicle {
	return mappers.DomainVehicleToResponseVehicle(domains)
}

func (h *VehicleHandler) mapRequestVehicleToDomainVehicle(vehicle requests.Vehicle) domain.Vehicle {
	return mappers.RequestVehicleToDomainVehicle(vehicle)
}
func (h *VehicleHandler) mapNewRequestVehicleToDomainVehicle(vehicle requests.Vehicle) domain.Vehicle {
	return mappers.NewRequestVehicleToDomainVehicle(vehicle)
}

func (h *VehicleHandler) mapDomainVehiclesToResponseVehicles(domains []domain.Vehicle) []responses.Vehicle {
	return mappers.DomainVehiclesToResponseVehicles(domains)
}
