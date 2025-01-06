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

type CreateVehicleHandlerTestSuite struct {
	suite.Suite
	domain        domains.Vehicle
	vehicle       requests.Vehicle
	vehicleEquals func(v domains.Vehicle) bool
}

func TestVehicleHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(CreateVehicleHandlerTestSuite))
}

func (suite *CreateVehicleHandlerTestSuite) SetupSubTest() {
	setup()
	router.HandleFunc("/vehicles/register", handler.HandleCreateVehicle).Methods(http.MethodPost)
	suite.vehicle = requests.Vehicle{
		Brand: "Toyota", Model: "Corolla", Year: 2020, Color: "Blue", Price: 20000, Status: "FOR_SALE", LicensePlate: "TEST-X",
	}
	suite.domain = domains.Vehicle{
		Brand:        suite.vehicle.Brand,
		Model:        suite.vehicle.Model,
		Year:         suite.vehicle.Year,
		Color:        suite.vehicle.Color,
		Price:        suite.vehicle.Price,
		Status:       suite.vehicle.Status,
		LicensePlate: suite.vehicle.LicensePlate,
	}
	suite.vehicleEquals = func(v domains.Vehicle) bool {
		return strings.EqualFold(v.Brand, suite.domain.Brand) &&
			strings.EqualFold(v.Model, suite.domain.Model) &&
			v.Year == suite.domain.Year &&
			strings.EqualFold(v.Color, suite.domain.Color) &&
			v.Price == suite.domain.Price &&
			suite.NotEmpty(v.ID)
	}
}

func (suite *CreateVehicleHandlerTestSuite) TearDownSubTest() {
	createVehicle.AssertExpectations(suite.T())
}

func (suite *CreateVehicleHandlerTestSuite) TestHandleCreateVehicle() {
	suite.Run("Should return Success", func() {

		createVehicle.On("Execute", mock.Anything, mock.MatchedBy(suite.vehicleEquals)).Return(nil)

		body, _ := json.Marshal(suite.vehicle)
		req := httptest.NewRequest(http.MethodPost, "/vehicles/register", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		suite.Equal(http.StatusCreated, rr.Code)
	})

	suite.Run("Should return BadRequest", func() {
		req := httptest.NewRequest(http.MethodPost, "/vehicles/register",
			bytes.NewBuffer([]byte("invalid body")))
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		suite.Equal(http.StatusBadRequest, rr.Code)
		var problemDetails http_errors.ProblemDetails
		err := json.NewDecoder(rr.Body).Decode(&problemDetails)
		suite.NoError(err)
		suite.Contains(problemDetails.Detail, "Failed to decode request body")
	})

	suite.Run("Should return BadRequest - invalid body", func() {
		suite.vehicle.Brand = ""
		body, _ := json.Marshal(suite.vehicle)
		req := httptest.NewRequest(http.MethodPost, "/vehicles/register", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		suite.Equal(http.StatusBadRequest, rr.Code)
		var problemDetails http_errors.ProblemDetails
		err := json.NewDecoder(rr.Body).Decode(&problemDetails)
		suite.NoError(err)
		suite.Contains(problemDetails.Detail, "Invalid request body")
	})

	suite.Run("Should return InternalServerError", func() {
		createVehicle.On("Execute", mock.Anything, mock.MatchedBy(suite.vehicleEquals)).
			Return(errors.New("some error"))

		body, _ := json.Marshal(suite.vehicle)
		req := httptest.NewRequest(http.MethodPost, "/vehicles/register", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		suite.Equal(http.StatusInternalServerError, rr.Code)
		var problemDetails http_errors.ProblemDetails
		err := json.NewDecoder(rr.Body).Decode(&problemDetails)
		suite.NoError(err)
		suite.Contains(problemDetails.Detail, "Failed to register vehicle")
	})
}
