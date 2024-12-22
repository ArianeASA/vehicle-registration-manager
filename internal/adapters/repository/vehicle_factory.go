package repository

import (
	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"log"
	"os"
	"vehicle-registration-manager/internal/adapters/repository/configs"
	"vehicle-registration-manager/internal/core/ports/out"
)

func VehicleRepositoryFactory(config *configs.DatabaseConfig) out.VehicleRepository {
	scope := os.Getenv("SCOPE")
	switch scope {
	case "prod":
		repo, err := NewVehicleRepository(config)
		if err != nil {
			log.Fatalf("failed to create SQL repository: %v", err)
		}
		return repo
	default:
		return NewLocalVehicleRepository()
	}
}
