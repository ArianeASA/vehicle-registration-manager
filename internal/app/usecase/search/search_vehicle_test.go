package search_test

import (
	"errors"
	"testing"
	"vehicle-registration-manager/internal/app/usecase/search"
	"vehicle-registration-manager/internal/core/domains"
	"vehicle-registration-manager/internal/core/ports/out"
	"vehicle-registration-manager/pkg/tracer"

	"github.com/stretchr/testify/suite"
)

type SearchVehicleTestSuite struct {
	suite.Suite
	repo      *out.VehicleRepositoryMock
	usecase   search.SearchVehicle
	tracer    *tracer.Tracer
	vehicle   domains.Vehicle
	vehicleID string
}

func TestSearchVehicleTestSuite(t *testing.T) {
	suite.Run(t, new(SearchVehicleTestSuite))
}

func (suite *SearchVehicleTestSuite) SetupSubTest() {
	suite.repo = new(out.VehicleRepositoryMock)
	suite.usecase = search.NewSearchVehicle(suite.repo)
	suite.tracer = tracer.NewFakeTracer()
	suite.vehicleID = "1"
	suite.vehicle = domains.Vehicle{
		ID:           suite.vehicleID,
		Brand:        "Toyota",
		Model:        "Corolla",
		Year:         2020,
		Color:        "Blue",
		Price:        20000,
		Status:       "FOR_SALE",
		LicensePlate: "TEST-X",
	}
}

func (suite *SearchVehicleTestSuite) TearDownSubTest() {
	suite.repo.AssertExpectations(suite.T())
}

func (suite *SearchVehicleTestSuite) TestExecute() {
	suite.Run("Should return vehicle on success", func() {
		suite.repo.On("FindByID", suite.tracer, suite.vehicleID).Return(suite.vehicle, nil)

		result, err := suite.usecase.Execute(suite.tracer, suite.vehicleID)
		suite.NoError(err)
		suite.Equal(suite.vehicle, result)
	})

	suite.Run("Should return error when repository fails", func() {
		suite.repo.On("FindByID", suite.tracer, suite.vehicleID).Return(domains.Vehicle{}, errors.New("some error"))

		result, err := suite.usecase.Execute(suite.tracer, suite.vehicleID)
		suite.Error(err)
		suite.Equal(domains.Vehicle{}, result)
	})

	suite.Run("Should return empty vehicle when not found", func() {
		suite.repo.On("FindByID", suite.tracer, suite.vehicleID).Return(domains.Vehicle{}, nil)

		result, err := suite.usecase.Execute(suite.tracer, suite.vehicleID)
		suite.NoError(err)
		suite.Equal(domains.Vehicle{}, result)
	})
}
