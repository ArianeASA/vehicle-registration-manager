package repository

import (
	"log"
	"os"
	"vehicle-registration-manager/internal/core/ports"
)

func NewVehicleRepository() ports.VehicleRepository {
	scope := os.Getenv("SCOPE")
	switch scope {
	case "prod":
		repo, err := NewSQLVehicleRepository("vehicles.db")
		if err != nil {
			log.Fatalf("failed to create SQL repository: %v", err)
		}
		return repo
	default:
		return NewMemoryVehicleRepository()

	}
}
