package ports

import "vehicle-registration-manager/internal/core/domain"

type VehicleRepository interface {
	Save(vehicle domain.Vehicle) error
	Update(vehicle domain.Vehicle) error
	FindAll() ([]domain.Vehicle, error)
}
