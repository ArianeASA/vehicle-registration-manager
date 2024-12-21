package usecase

import (
	"vehicle-registration-manager/internal/core/domains"
	"vehicle-registration-manager/internal/core/ports/out"
	"vehicle-registration-manager/pkg/tracer"
)

type ListVehicles struct {
	repo out.VehicleRepository
}

func NewListVehicles(repo out.VehicleRepository) *ListVehicles {
	return &ListVehicles{repo: repo}
}

func (uc *ListVehicles) Execute(trc *tracer.Tracer) ([]domains.Vehicle, error) {
	trc.Logger.Info("Init list vehicles")
	return uc.repo.FindAll(trc)
}
