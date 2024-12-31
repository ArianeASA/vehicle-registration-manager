package list_test

import (
	"errors"
	"testing"
	"vehicle-registration-manager/internal/app/usecase/list"
	"vehicle-registration-manager/internal/core/domains"
	"vehicle-registration-manager/internal/core/ports/out"
	"vehicle-registration-manager/pkg/tracer"

	"github.com/stretchr/testify/suite"
)

type ListVehiclesTestSuite struct {
	suite.Suite
	repo     *out.VehicleRepositoryMock
	usecase  list.ListVehicles
	tracer   *tracer.Tracer
	vehicles []domains.Vehicle
}

func TestListVehiclesTestSuite(t *testing.T) {
	suite.Run(t, new(ListVehiclesTestSuite))
}

func (suite *ListVehiclesTestSuite) SetupSubTest() {
	suite.repo = new(out.VehicleRepositoryMock)
	suite.usecase = list.NewListVehicles(suite.repo)
	suite.tracer = tracer.NewFakeTracer()
	suite.vehicles = []domains.Vehicle{
		{ID: "1", Brand: "Toyota", Model: "Corolla", Year: 2020, Color: "Blue", Price: 20000},
		{ID: "2", Brand: "Honda", Model: "Civic", Year: 2019, Color: "Red", Price: 18000},
	}
}

func (suite *ListVehiclesTestSuite) TearDownSubTest() {
	suite.repo.AssertExpectations(suite.T())
}

func (suite *ListVehiclesTestSuite) TestExecute() {
	suite.Run("Should return vehicles on success", func() {
		suite.repo.On("FindAll", suite.tracer).Return(suite.vehicles, nil)

		result, err := suite.usecase.Execute(suite.tracer)
		suite.NoError(err)
		suite.Equal(suite.vehicles, result)
	})

	suite.Run("Should return error when repository fails", func() {
		suite.repo.On("FindAll", suite.tracer).Return([]domains.Vehicle{}, errors.New("some error"))

		result, err := suite.usecase.Execute(suite.tracer)
		suite.Error(err)
		suite.Empty(result)
	})
}
