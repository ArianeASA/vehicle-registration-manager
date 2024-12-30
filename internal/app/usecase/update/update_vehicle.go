package update

import (
	"vehicle-registration-manager/internal/core/domains"
	"vehicle-registration-manager/internal/core/ports/out"
	"vehicle-registration-manager/pkg/tracer"
)

type (
	updateVehicle struct {
		repo out.VehicleRepository
	}

	UpdateVehicle interface {
		Execute(tcr *tracer.Tracer, vehicle domains.Vehicle) error
	}
)

func NewUpdateVehicle(repo out.VehicleRepository) UpdateVehicle {
	return &updateVehicle{repo: repo}
}

func (uc *updateVehicle) Execute(tcr *tracer.Tracer, vehicle domains.Vehicle) error {
	tcr.Logger.Info("Init update vehicle")
	vehicleDB, err := uc.repo.FindByID(tcr, vehicle.ID)
	if err != nil {
		return err
	}

	if !vehicleDB.Exist() {
		tcr.Logger.Error("Vehicle not found", domains.ErrVehicleNotFound)
		return domains.ErrVehicleNotFound
	}

	return uc.repo.Update(tcr, vehicle)
}
