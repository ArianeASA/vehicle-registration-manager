package create

import (
	"github.com/stretchr/testify/mock"
	"vehicle-registration-manager/internal/core/domains"
	"vehicle-registration-manager/pkg/tracer"
)

// MockCreateVehicle is a mock implementation of the CreateVehiclesUseCase interface
type MockCreateVehicle struct {
	mock.Mock
}

func (m *MockCreateVehicle) Execute(trc *tracer.Tracer, vehicle domains.Vehicle) error {
	args := m.Called(trc, vehicle)
	return args.Error(0)
}
