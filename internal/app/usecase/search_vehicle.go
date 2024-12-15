package usecase

import (
	"vehicle-registration-manager/internal/core/domain"
	"vehicle-registration-manager/internal/core/ports"
)

type SearchVehicle struct {
	repo ports.VehicleRepository
}

func NewSearchVehicle(repo ports.VehicleRepository) *SearchVehicle {
	return &SearchVehicle{repo: repo}
}

func (uc *SearchVehicle) Execute(id string) (domain.Vehicle, error) {
	return uc.repo.FindByID(id)
}
