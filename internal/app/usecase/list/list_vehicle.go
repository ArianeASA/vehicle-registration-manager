package list

import (
	"vehicle-registration-manager/internal/core/domains"
	"vehicle-registration-manager/internal/core/ports/out"
	"vehicle-registration-manager/pkg/tracer"
)

type (
	listVehicles struct {
		repo out.VehicleRepository
	}

	ListVehicles interface {
		Execute(trc *tracer.Tracer) ([]domains.Vehicle, error)
	}
)

func NewListVehicles(repo out.VehicleRepository) ListVehicles {
	return &listVehicles{repo: repo}
}

func (uc *listVehicles) Execute(trc *tracer.Tracer) ([]domains.Vehicle, error) {
	trc.Logger.Info("Init list vehicles")
	return uc.repo.FindAll(trc)
}
