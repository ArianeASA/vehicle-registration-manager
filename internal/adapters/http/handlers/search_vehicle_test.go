package handlers_test

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"vehicle-registration-manager/internal/core/domains"
	"vehicle-registration-manager/pkg/http_errors"
)

type SearchVehicleHandlerTestSuite struct {
	suite.Suite
	domain        domains.Vehicle
	vehicleID     string
	vehicleEquals func(v domains.Vehicle) bool
}

func TestSearchVehicleHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(SearchVehicleHandlerTestSuite))
}

func (suite *SearchVehicleHandlerTestSuite) SetupSubTest() {
	setup()
	router.StrictSlash(true)
	router.HandleFunc("/vehicles/", handler.HandleSearchVehicleByID).Methods(http.MethodGet)
	router.HandleFunc("/vehicles/{id}", handler.HandleSearchVehicleByID).
		Methods(http.MethodGet)
	suite.vehicleID = "1"
	suite.domain = domains.Vehicle{
		ID:    suite.vehicleID,
		Brand: "Toyota",
		Model: "Corolla",
		Year:  2020,
		Color: "Blue",
		Price: 20000,
	}
	suite.vehicleEquals = func(v domains.Vehicle) bool {
		return strings.EqualFold(v.ID, suite.domain.ID) &&
			strings.EqualFold(v.Brand, suite.domain.Brand) &&
			strings.EqualFold(v.Model, suite.domain.Model) &&
			v.Year == suite.domain.Year &&
			strings.EqualFold(v.Color, suite.domain.Color) &&
			v.Price == suite.domain.Price
	}
}

func (suite *SearchVehicleHandlerTestSuite) TearDownSubTest() {
	searchVehicle.AssertExpectations(suite.T())
}

func (suite *SearchVehicleHandlerTestSuite) TestHandleSearchVehicleByID() {
	suite.Run("Should return Success", func() {
		searchVehicle.On("Execute", mock.Anything, suite.vehicleID).Return(suite.domain, nil)

		req := httptest.NewRequest(http.MethodGet, "/vehicles/1", nil)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		suite.Equal(http.StatusOK, rr.Code)
		var response domains.Vehicle
		err := json.NewDecoder(rr.Body).Decode(&response)
		suite.NoError(err)
		suite.True(suite.vehicleEquals(response))
	})

	suite.Run("Should return BadRequest when id is empty", func() {
		req := httptest.NewRequest(http.MethodGet, "/vehicles/", nil)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		suite.Equal(http.StatusBadRequest, rr.Code)
		var problemDetails http_errors.ProblemDetails
		err := json.NewDecoder(rr.Body).Decode(&problemDetails)
		suite.NoError(err)
		suite.Contains(problemDetails.Detail, "Invalid id")
	})

	suite.Run("Should return InternalServerError", func() {
		searchVehicle.On("Execute", mock.Anything, suite.vehicleID).Return(domains.Vehicle{}, errors.New("some error"))

		req := httptest.NewRequest(http.MethodGet, "/vehicles/1", nil)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		suite.Equal(http.StatusInternalServerError, rr.Code)
		var problemDetails http_errors.ProblemDetails
		err := json.NewDecoder(rr.Body).Decode(&problemDetails)
		suite.NoError(err)
		suite.Contains(problemDetails.Detail, "Failed to search vehicles")
	})

	suite.Run("Should return NotFound", func() {
		searchVehicle.On("Execute", mock.Anything, suite.vehicleID).Return(domains.Vehicle{}, nil)

		req := httptest.NewRequest(http.MethodGet, "/vehicles/1", nil)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		suite.Equal(http.StatusNotFound, rr.Code)
		var problemDetails http_errors.ProblemDetails
		err := json.NewDecoder(rr.Body).Decode(&problemDetails)
		suite.NoError(err)
		suite.Contains(problemDetails.Detail, "Vehicle not found")
	})
}
