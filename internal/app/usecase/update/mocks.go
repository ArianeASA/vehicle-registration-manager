package update

import (
	"github.com/stretchr/testify/mock"
	"vehicle-registration-manager/internal/core/domains"
	"vehicle-registration-manager/pkg/tracer"
)

// MockUpdateVehicle is a mock implementation of the UpdateVehicleUseCase interface
type MockUpdateVehicle struct {
	mock.Mock
}

func (u *MockUpdateVehicle) Execute(trc *tracer.Tracer, vehicle domains.Vehicle) error {
	args := u.Called(trc, vehicle)
	return args.Error(0)
}
