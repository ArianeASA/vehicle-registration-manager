package mappers

import (
	"vehicle-registration-manager/internal/adapters/repository/entities"
	"vehicle-registration-manager/internal/core/domains"
)

func DomainToEntity(vehicle domains.Vehicle) entities.Vehicle {
	return entities.Vehicle{
		ID:    vehicle.ID,
		Brand: vehicle.Brand,
		Model: vehicle.Model,
		Year:  vehicle.Year,
		Color: vehicle.Color,
		Price: vehicle.Price,
	}
}

func EntityToDomain(entity entities.Vehicle) domains.Vehicle {
	return domains.Vehicle{
		ID:    entity.ID,
		Brand: entity.Brand,
		Model: entity.Model,
		Year:  entity.Year,
		Color: entity.Color,
		Price: entity.Price,
	}
}
