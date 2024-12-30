package create

import (
	"vehicle-registration-manager/internal/core/domains"
	"vehicle-registration-manager/internal/core/ports/out"
	"vehicle-registration-manager/pkg/tracer"
)

type (
	createVehicle struct {
		repo out.VehicleRepository
	}

	CreateVehicle interface {
		Execute(trc *tracer.Tracer, vehicle domains.Vehicle) error
	}
)

func NewCreateVehicle(repo out.VehicleRepository) CreateVehicle {
	return &createVehicle{repo: repo}
}

func (uc *createVehicle) Execute(trc *tracer.Tracer, vehicle domains.Vehicle) error {
	trc.Logger.Info("Register vehicle")
	return uc.repo.Save(trc, vehicle)
}
