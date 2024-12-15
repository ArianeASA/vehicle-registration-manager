package usecase

import (
	"vehicle-registration-manager/internal/core/domain"
	"vehicle-registration-manager/internal/core/ports"
	"vehicle-registration-manager/pkg/tracer"
)

type ListVehicles struct {
	repo ports.VehicleRepository
}

func NewListVehicles(repo ports.VehicleRepository) *ListVehicles {
	return &ListVehicles{repo: repo}
}

func (uc *ListVehicles) Execute(trc *tracer.Tracer) ([]domain.Vehicle, error) {
	trc.Logger.Info("ListVehicles.Execute")
	return uc.repo.FindAll()
}
