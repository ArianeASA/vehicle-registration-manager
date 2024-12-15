package usecase

import (
	"vehicle-registration-manager/internal/core/domain"
	"vehicle-registration-manager/internal/core/ports"
	"vehicle-registration-manager/pkg/tracer"
)

type RegisterVehicle struct {
	repo ports.VehicleRepository
}

func NewRegisterVehicle(repo ports.VehicleRepository) *RegisterVehicle {
	return &RegisterVehicle{repo: repo}
}

func (uc *RegisterVehicle) Execute(trc *tracer.Tracer, vehicle domain.Vehicle) error {
	trc.Logger.Info("Register vehicle")
	return uc.repo.Save(vehicle)
}
