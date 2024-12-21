package usecase

import (
	"vehicle-registration-manager/internal/core/domains"
	"vehicle-registration-manager/internal/core/ports/out"
	"vehicle-registration-manager/pkg/tracer"
)

type RegisterVehicle struct {
	repo out.VehicleRepository
}

func NewRegisterVehicle(repo out.VehicleRepository) *RegisterVehicle {
	return &RegisterVehicle{repo: repo}
}

func (uc *RegisterVehicle) Execute(trc *tracer.Tracer, vehicle domains.Vehicle) error {
	trc.Logger.Info("Register vehicle")
	return uc.repo.Save(trc, vehicle)
}
