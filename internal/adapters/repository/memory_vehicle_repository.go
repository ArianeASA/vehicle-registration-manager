package repository

import (
	"fmt"
	"vehicle-registration-manager/internal/core/domain"

	"sync"
)

type MemoryVehicleRepository struct {
	vehicles map[string]domain.Vehicle
	mu       sync.RWMutex
}

func NewMemoryVehicleRepository() *MemoryVehicleRepository {
	return &MemoryVehicleRepository{
		vehicles: make(map[string]domain.Vehicle),
	}
}

func (r *MemoryVehicleRepository) Save(vehicle domain.Vehicle) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.vehicles[vehicle.ID] = vehicle
	return nil
}

func (r *MemoryVehicleRepository) Update(vehicle domain.Vehicle) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.vehicles[vehicle.ID]; !exists {
		return fmt.Errorf("vehicle not found")
	}
	r.vehicles[vehicle.ID] = vehicle
	return nil
}

func (r *MemoryVehicleRepository) FindAll() ([]domain.Vehicle, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	vehicles := make([]domain.Vehicle, 0, len(r.vehicles))
	for _, v := range r.vehicles {
		vehicles = append(vehicles, v)
	}
	return vehicles, nil
}
