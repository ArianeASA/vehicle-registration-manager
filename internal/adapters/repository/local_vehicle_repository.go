package repository

import (
	"fmt"
	"sync"
	"vehicle-registration-manager/internal/core/domains"
	"vehicle-registration-manager/pkg/tracer"
)

type LocalVehicleRepository struct {
	vehicles map[string]domains.Vehicle
	mu       sync.RWMutex
}

func NewLocalVehicleRepository() *LocalVehicleRepository {
	return &LocalVehicleRepository{
		vehicles: make(map[string]domains.Vehicle),
	}
}

func (r *LocalVehicleRepository) Save(_ *tracer.Tracer, vehicle domains.Vehicle) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.vehicles[vehicle.ID] = vehicle
	return nil
}

func (r *LocalVehicleRepository) Update(_ *tracer.Tracer, vehicle domains.Vehicle) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.vehicles[vehicle.ID]; !exists {
		return fmt.Errorf("vehicle not found")
	}
	r.vehicles[vehicle.ID] = vehicle
	return nil
}

func (r *LocalVehicleRepository) FindAll(_ *tracer.Tracer) ([]domains.Vehicle, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	vehicles := make([]domains.Vehicle, 0, len(r.vehicles))
	for _, v := range r.vehicles {
		vehicles = append(vehicles, v)
	}
	return vehicles, nil
}

func (r *LocalVehicleRepository) FindByID(_ *tracer.Tracer, id string) (domains.Vehicle, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.vehicles[id], nil
}
