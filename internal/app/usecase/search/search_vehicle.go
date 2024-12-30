package search

import (
	"vehicle-registration-manager/internal/core/domains"
	"vehicle-registration-manager/internal/core/ports/out"
	"vehicle-registration-manager/pkg/tracer"
)

type (
	searchVehicle struct {
		repo out.VehicleRepository
	}

	SearchVehicle interface {
		Execute(tcr *tracer.Tracer, id string) (domains.Vehicle, error)
	}
)

func NewSearchVehicle(repo out.VehicleRepository) SearchVehicle {
	return &searchVehicle{repo: repo}
}

func (uc *searchVehicle) Execute(tcr *tracer.Tracer, id string) (domains.Vehicle, error) {
	tcr.Logger.Info("Init search vehicle")
	return uc.repo.FindByID(tcr, id)
}
