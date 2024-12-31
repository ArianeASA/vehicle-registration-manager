package repository

import (
	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"os"
	"vehicle-registration-manager/internal/adapters/repository/configs"
	"vehicle-registration-manager/internal/core/ports/out"
)

func VehicleRepositoryFactory(config configs.DatabaseConfigs) out.VehicleRepository {
	scope := os.Getenv("SCOPE")
	switch scope {
	case "prod":
		return NewVehicleRepository(config)
	default:
		return NewLocalVehicleRepository()
	}
}
