package usecase

import (
	"vehicle-registration-manager/internal/core/domains"
	"vehicle-registration-manager/internal/core/ports/out"
	"vehicle-registration-manager/pkg/tracer"
)

type SearchVehicle struct {
	repo out.VehicleRepository
}

func NewSearchVehicle(repo out.VehicleRepository) *SearchVehicle {
	return &SearchVehicle{repo: repo}
}

func (uc *SearchVehicle) Execute(tcr *tracer.Tracer, id string) (domains.Vehicle, error) {
	tcr.Logger.Info("Init search vehicle")
	return uc.repo.FindByID(tcr, id)
}
