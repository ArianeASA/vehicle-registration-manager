package out

import (
	"github.com/stretchr/testify/mock"
	"vehicle-registration-manager/internal/core/domains"
	"vehicle-registration-manager/pkg/tracer"
)

type VehicleRepository interface {
	Save(tcr *tracer.Tracer, vehicle domains.Vehicle) error
	Update(tcr *tracer.Tracer, vehicle domains.Vehicle) error
	FindAll(tcr *tracer.Tracer) ([]domains.Vehicle, error)
	FindByID(tcr *tracer.Tracer, id string) (domains.Vehicle, error)
}

// VehicleRepositoryMock is a mock of VehicleRepository interface
type VehicleRepositoryMock struct {
	mock.Mock
}

func (m *VehicleRepositoryMock) Save(tcr *tracer.Tracer, vehicle domains.Vehicle) error {
	args := m.Called(tcr, vehicle)
	return args.Error(0)
}

func (m *VehicleRepositoryMock) Update(tcr *tracer.Tracer, vehicle domains.Vehicle) error {
	args := m.Called(tcr, vehicle)
	return args.Error(0)
}

func (m *VehicleRepositoryMock) FindAll(tcr *tracer.Tracer) ([]domains.Vehicle, error) {
	args := m.Called(tcr)
	return args.Get(0).([]domains.Vehicle), args.Error(1)
}

func (m *VehicleRepositoryMock) FindByID(tcr *tracer.Tracer, id string) (domains.Vehicle, error) {
	args := m.Called(tcr, id)
	return args.Get(0).(domains.Vehicle), args.Error(1)

}
