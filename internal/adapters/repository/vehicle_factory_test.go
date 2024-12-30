package repository_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"vehicle-registration-manager/internal/adapters/repository"
	"vehicle-registration-manager/internal/adapters/repository/configs"
)

func TestVehicleRepositoryFactory_Local(t *testing.T) {
	mockDB := new(configs.MockDatabaseConfig)
	db, _, _ := sqlmock.New()
	mockDB.On("GetDB").Return(db)

	repo := repository.VehicleRepositoryFactory(mockDB)

	assert.NotNil(t, repo)
}

func TestVehicleRepositoryFactory_Prod(t *testing.T) {
	os.Setenv("SCOPE", "prod")
	mockDB := new(configs.MockDatabaseConfig)
	db, _, _ := sqlmock.New()
	mockDB.On("GetDB").Return(db)

	repo := repository.VehicleRepositoryFactory(mockDB)

	assert.NotNil(t, repo)
}
