package usecase

import (
	"vehicle-registration-manager/internal/core/domain"
	"vehicle-registration-manager/internal/core/ports"
)

type ListVehicles struct {
	repo ports.VehicleRepository
}

func NewListVehicles(repo ports.VehicleRepository) *ListVehicles {
	return &ListVehicles{repo: repo}
}

func (uc *ListVehicles) Execute() ([]domain.Vehicle, error) {
	return uc.repo.FindAll()
}
