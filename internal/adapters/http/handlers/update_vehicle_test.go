package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"vehicle-registration-manager/internal/adapters/http/requests"
	"vehicle-registration-manager/internal/core/domains"
	"vehicle-registration-manager/pkg/http_errors"
)

type UpdateVehicleHandlerTestSuite struct {
	suite.Suite
	domain        domains.Vehicle
	vehicle       requests.Vehicle
	vehicleEquals func(v domains.Vehicle) bool
}

func TestUpdateVehicleHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(UpdateVehicleHandlerTestSuite))
}

func (suite *UpdateVehicleHandlerTestSuite) SetupSubTest() {
	setup()
	router.StrictSlash(true)
	router.HandleFunc("/vehicles/", handler.HandleUpdateVehicle).Methods(http.MethodPut)
	router.HandleFunc("/vehicles/{id}", handler.HandleUpdateVehicle).
		Methods(http.MethodPut)

	suite.vehicle = requests.Vehicle{
		Brand: "Toyota", Model: "Corolla", Year: 2020, Color: "Blue", Price: 20000,
	}
	suite.domain = domains.Vehicle{
		ID:    "1",
		Brand: suite.vehicle.Brand,
		Model: suite.vehicle.Model,
		Year:  suite.vehicle.Year,
		Color: suite.vehicle.Color,
		Price: suite.vehicle.Price,
	}
	suite.vehicleEquals = func(v domains.Vehicle) bool {
		return strings.EqualFold(v.Brand, suite.domain.Brand) &&
			strings.EqualFold(v.Model, suite.domain.Model) &&
			v.Year == suite.domain.Year &&
			strings.EqualFold(v.Color, suite.domain.Color) &&
			v.Price == suite.domain.Price &&
			strings.EqualFold(v.ID, suite.domain.ID)
	}
}

func (suite *UpdateVehicleHandlerTestSuite) TearDownSubTest() {
	updateVehicle.AssertExpectations(suite.T())
}

func (suite *UpdateVehicleHandlerTestSuite) TestHandleUpdateVehicle() {
	suite.Run("Should return Success", func() {
		updateVehicle.On("Execute", mock.Anything, mock.MatchedBy(suite.vehicleEquals)).
			Return(nil)

		body, _ := json.Marshal(suite.vehicle)
		req := httptest.NewRequest(http.MethodPut, "/vehicles/1", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		suite.Equal(http.StatusOK, rr.Code)
	})

	suite.Run("Should return BadRequest when body is invalid", func() {
		req := httptest.NewRequest(http.MethodPut, "/vehicles/1", bytes.NewBuffer([]byte("invalid body")))
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		suite.Equal(http.StatusBadRequest, rr.Code)
		var problemDetails http_errors.ProblemDetails
		err := json.NewDecoder(rr.Body).Decode(&problemDetails)
		suite.NoError(err)
		suite.Contains(problemDetails.Detail, "Failed to decode request body")
	})

	suite.Run("Should return StatusBadRequest when id is empty", func() {
		body, _ := json.Marshal(suite.vehicle)
		req := httptest.NewRequest(http.MethodPut, "/vehicles/", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		suite.Equal(http.StatusBadRequest, rr.Code)
		var problemDetails http_errors.ProblemDetails
		err := json.NewDecoder(rr.Body).Decode(&problemDetails)
		suite.NoError(err)
		suite.Contains(problemDetails.Detail, "Invalid id")
	})

	suite.Run("Should return InternalServerError", func() {
		updateVehicle.On("Execute", mock.Anything, mock.MatchedBy(suite.vehicleEquals)).
			Return(errors.New("some error"))

		body, _ := json.Marshal(suite.vehicle)
		req := httptest.NewRequest(http.MethodPut, "/vehicles/1", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		suite.Equal(http.StatusInternalServerError, rr.Code)
		var problemDetails http_errors.ProblemDetails
		err := json.NewDecoder(rr.Body).Decode(&problemDetails)
		suite.NoError(err)
		suite.Contains(problemDetails.Detail, "Failed to update vehicle")
	})

	suite.Run("Should return NotFound", func() {
		updateVehicle.On("Execute", mock.Anything, mock.MatchedBy(suite.vehicleEquals)).Return(domains.ErrVehicleNotFound)

		body, _ := json.Marshal(suite.vehicle)
		req := httptest.NewRequest(http.MethodPut, "/vehicles/1", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		suite.Equal(http.StatusNotFound, rr.Code)
		var problemDetails http_errors.ProblemDetails
		err := json.NewDecoder(rr.Body).Decode(&problemDetails)
		suite.NoError(err)
		suite.Contains(problemDetails.Detail, "Vehicle not found")
	})
}
