package search

import (
	"github.com/stretchr/testify/mock"
	"vehicle-registration-manager/internal/core/domains"
	"vehicle-registration-manager/pkg/tracer"
)

// MockSearchVehicle is a mock implementation of the SearchVehicleUseCase interface
type MockSearchVehicle struct {
	mock.Mock
}

func (s *MockSearchVehicle) Execute(trc *tracer.Tracer, id string) (domains.Vehicle, error) {
	args := s.Called(trc, id)
	return args.Get(0).(domains.Vehicle), args.Error(1)
}
