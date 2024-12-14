package usecase

import (
	"vehicle-registration-manager/internal/core/domain"
	"vehicle-registration-manager/internal/core/ports"
)

type RegisterVehicle struct {
	repo ports.VehicleRepository
}

func NewRegisterVehicle(repo ports.VehicleRepository) *RegisterVehicle {
	return &RegisterVehicle{repo: repo}
}

func (uc *RegisterVehicle) Execute(vehicle domain.Vehicle) error {
	return uc.repo.Save(vehicle)
}
