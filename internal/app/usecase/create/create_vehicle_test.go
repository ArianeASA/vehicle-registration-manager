package create_test

import (
	"errors"
	"testing"
	"vehicle-registration-manager/internal/app/usecase/create"
	"vehicle-registration-manager/internal/core/domains"
	"vehicle-registration-manager/internal/core/ports/out"
	"vehicle-registration-manager/pkg/tracer"

	"github.com/stretchr/testify/suite"
)

type CreateVehicleTestSuite struct {
	suite.Suite
	repo    *out.VehicleRepositoryMock
	usecase create.CreateVehicle
	tracer  *tracer.Tracer
	vehicle domains.Vehicle
}

func TestCreateVehicleTestSuite(t *testing.T) {
	suite.Run(t, new(CreateVehicleTestSuite))
}

func (suite *CreateVehicleTestSuite) SetupSubTest() {
	suite.repo = new(out.VehicleRepositoryMock)
	suite.usecase = create.NewCreateVehicle(suite.repo)
	suite.tracer = tracer.NewFakeTracer()
	suite.vehicle = domains.Vehicle{
		ID:           "1",
		Brand:        "Toyota",
		Model:        "Corolla",
		Year:         2020,
		Color:        "Blue",
		Price:        20000,
		Status:       "FOR_SALE",
		LicensePlate: "TEST-X",
	}
}

func (suite *CreateVehicleTestSuite) TearDownSubTest() {
	suite.repo.AssertExpectations(suite.T())
}

func (suite *CreateVehicleTestSuite) TestExecute() {
	suite.Run("Should return nil on success", func() {
		suite.repo.On("Save", suite.tracer, suite.vehicle).Return(nil)

		err := suite.usecase.Execute(suite.tracer, suite.vehicle)
		suite.NoError(err)
	})

	suite.Run("Should return error when repository fails", func() {
		suite.repo.On("Save", suite.tracer, suite.vehicle).Return(errors.New("some error"))

		err := suite.usecase.Execute(suite.tracer, suite.vehicle)
		suite.Error(err)
	})
}
