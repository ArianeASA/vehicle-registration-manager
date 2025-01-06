package update_test

import (
	"errors"
	"testing"
	"vehicle-registration-manager/internal/app/usecase/update"
	"vehicle-registration-manager/internal/core/domains"
	"vehicle-registration-manager/internal/core/domains/consts"
	"vehicle-registration-manager/internal/core/ports/out"
	"vehicle-registration-manager/pkg/tracer"

	"github.com/stretchr/testify/suite"
)

type UpdateVehicleTestSuite struct {
	suite.Suite
	repo             *out.VehicleRepositoryMock
	usecase          update.UpdateVehicle
	tracer           *tracer.Tracer
	vehicle          domains.Vehicle
	newVehicleStatus domains.Vehicle
}

func TestUpdateVehicleTestSuite(t *testing.T) {
	suite.Run(t, new(UpdateVehicleTestSuite))
}

func (suite *UpdateVehicleTestSuite) SetupSubTest() {
	suite.repo = new(out.VehicleRepositoryMock)
	suite.usecase = update.NewUpdateVehicle(suite.repo)
	suite.tracer = tracer.NewFakeTracer()
	suite.vehicle = domains.Vehicle{
		ID:     "1",
		Brand:  "Toyota",
		Model:  "Corolla",
		Year:   2020,
		Color:  "Blue",
		Price:  20000,
		Status: consts.StatusForSale,
	}
	suite.newVehicleStatus = suite.vehicle
	suite.newVehicleStatus.Status = consts.StatusReserved
}

func (suite *UpdateVehicleTestSuite) TearDownSubTest() {
	suite.repo.AssertExpectations(suite.T())
}

func (suite *UpdateVehicleTestSuite) TestExecute() {
	suite.Run("Should return nil on success", func() {
		suite.repo.On("FindByID", suite.tracer, suite.vehicle.ID).Return(suite.vehicle, nil)
		suite.repo.On("Update", suite.tracer, suite.vehicle).Return(nil)

		err := suite.usecase.Execute(suite.tracer, suite.vehicle)
		suite.NoError(err)
	})

	suite.Run("Should return nil on success- update status", func() {
		suite.repo.On("FindByID", suite.tracer, suite.vehicle.ID).Return(suite.vehicle, nil)
		suite.repo.On("Update", suite.tracer, suite.newVehicleStatus).Return(nil)

		err := suite.usecase.Execute(suite.tracer, suite.newVehicleStatus)
		suite.NoError(err)
	})

	suite.Run("Should return error when update status invalid", func() {
		suite.newVehicleStatus.Status = "invalid"
		expectedError := errors.New("invalid status transition from FOR_SALE to invalid. " +
			"Rule status permitted: from FOR_SALE to [RESERVED CANCELED]")
		suite.repo.On("FindByID", suite.tracer, suite.vehicle.ID).Return(suite.vehicle, nil)

		err := suite.usecase.Execute(suite.tracer, suite.newVehicleStatus)
		suite.Error(err)
		suite.Equal(expectedError, err)
	})

	suite.Run("Should return error when vehicle not found", func() {
		suite.repo.On("FindByID", suite.tracer, suite.vehicle.ID).Return(domains.Vehicle{}, nil)

		err := suite.usecase.Execute(suite.tracer, suite.vehicle)
		suite.Error(err)
		suite.Equal(domains.ErrVehicleNotFound, err)
	})

	suite.Run("Should return error when repository fails to find vehicle", func() {
		expectedError := errors.New("some error")
		suite.repo.On("FindByID", suite.tracer, suite.vehicle.ID).Return(domains.Vehicle{}, expectedError)

		err := suite.usecase.Execute(suite.tracer, suite.vehicle)
		suite.Error(err)
		suite.Equal(expectedError, err)
	})

	suite.Run("Should return error when repository fails to update vehicle", func() {
		expectedError := errors.New("some error")
		suite.repo.On("FindByID", suite.tracer, suite.vehicle.ID).Return(suite.vehicle, nil)
		suite.repo.On("Update", suite.tracer, suite.vehicle).Return(expectedError)

		err := suite.usecase.Execute(suite.tracer, suite.vehicle)
		suite.Error(err)
		suite.Equal(expectedError, err)
	})
}
