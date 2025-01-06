package mappers

import (
	"github.com/google/uuid"
	"vehicle-registration-manager/internal/adapters/http/requests"
	"vehicle-registration-manager/internal/core/domains"
	"vehicle-registration-manager/internal/core/domains/consts"
)

func RequestVehicleToDomainVehicle(vehicle requests.Vehicle) domains.Vehicle {
	return domains.Vehicle{
		Brand:  vehicle.Brand,
		Model:  vehicle.Model,
		Year:   vehicle.Year,
		Color:  vehicle.Color,
		Price:  vehicle.Price,
		Status: vehicle.Status,
	}
}

func NewRequestVehicleToDomainVehicle(vehicle requests.Vehicle) domains.Vehicle {

	uuidValue, _ := uuid.NewV7()
	return domains.Vehicle{
		ID:           uuidValue.String(),
		Brand:        vehicle.Brand,
		Model:        vehicle.Model,
		Year:         vehicle.Year,
		Color:        vehicle.Color,
		Price:        vehicle.Price,
		Status:       consts.StatusForSale,
		LicensePlate: vehicle.LicensePlate,
	}
}
