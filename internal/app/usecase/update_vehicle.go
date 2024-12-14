package usecase

import (
	"vehicle-registration-manager/internal/core/domain"
	"vehicle-registration-manager/internal/core/ports"
)

type UpdateVehicle struct {
	repo ports.VehicleRepository
}

func NewUpdateVehicle(repo ports.VehicleRepository) *UpdateVehicle {
	return &UpdateVehicle{repo: repo}
}

func (uc *UpdateVehicle) Execute(vehicle domain.Vehicle) error {
	return uc.repo.Update(vehicle)
}
