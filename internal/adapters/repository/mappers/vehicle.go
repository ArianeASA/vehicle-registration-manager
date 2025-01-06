package mappers

import (
	"vehicle-registration-manager/internal/adapters/repository/entities"
	"vehicle-registration-manager/internal/core/domains"
	"vehicle-registration-manager/internal/core/domains/consts"
)

func EntityToDomain(entity entities.Vehicle) domains.Vehicle {
	return domains.Vehicle{
		ID:           entity.ID,
		Brand:        entity.Brand,
		Model:        entity.Model,
		Year:         entity.Year,
		Color:        entity.Color,
		Price:        entity.Price,
		LicensePlate: entity.LicensePlate,
		Status:       consts.Status(entity.Status),
	}
}
