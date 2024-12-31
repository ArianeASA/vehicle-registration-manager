package list

import (
	"github.com/stretchr/testify/mock"
	"vehicle-registration-manager/internal/core/domains"
	"vehicle-registration-manager/pkg/tracer"
)

// MockListVehicles is a mock implementation of the ListVehiclesUseCase interface
type MockListVehicles struct {
	mock.Mock
}

func (l *MockListVehicles) Execute(trc *tracer.Tracer) ([]domains.Vehicle, error) {
	args := l.Called(trc)
	return args.Get(0).([]domains.Vehicle), args.Error(1)
}
