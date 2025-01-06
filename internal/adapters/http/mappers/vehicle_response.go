package mappers

import (
	"vehicle-registration-manager/internal/adapters/http/responses"
	"vehicle-registration-manager/internal/core/domains"
)

func DomainVehicleToResponseVehicle(vehicle domains.Vehicle) responses.Vehicle {
	return responses.Vehicle{
		Id:           vehicle.ID,
		Brand:        vehicle.Brand,
		Model:        vehicle.Model,
		Year:         vehicle.Year,
		Color:        vehicle.Color,
		Price:        vehicle.Price,
		LicensePlate: vehicle.LicensePlate,
		Status:       vehicle.Status,
	}
}

func DomainVehiclesToResponseVehicles(vehicles []domains.Vehicle) []responses.Vehicle {
	var responseVehicles []responses.Vehicle
	for _, vehicle := range vehicles {
		responseVehicles = append(responseVehicles, DomainVehicleToResponseVehicle(vehicle))
	}
	return responseVehicles
}
