package mappers

import (
	"github.com/google/uuid"
	"vehicle-registration-manager/internal/adapters/http/requests"
	"vehicle-registration-manager/internal/core/domain"
)

func RequestVehicleToDomainVehicle(vehicle requests.Vehicle) domain.Vehicle {
	return domain.Vehicle{
		Brand: vehicle.Brand,
		Model: vehicle.Model,
		Year:  vehicle.Year,
		Color: vehicle.Color,
		Price: vehicle.Price,
	}
}

func NewRequestVehicleToDomainVehicle(vehicle requests.Vehicle) domain.Vehicle {

	uuidValue, _ := uuid.NewV7()
	return domain.Vehicle{
		ID:    uuidValue.String(),
		Brand: vehicle.Brand,
		Model: vehicle.Model,
		Year:  vehicle.Year,
		Color: vehicle.Color,
		Price: vehicle.Price,
	}
}
