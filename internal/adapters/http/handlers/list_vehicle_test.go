package handlers_test

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
	"vehicle-registration-manager/internal/adapters/http/responses"
	"vehicle-registration-manager/internal/core/domains"
	"vehicle-registration-manager/pkg/http_errors"
)

type ListVehicleHandlerTestSuite struct {
	suite.Suite
}

func TestListVehicleHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(ListVehicleHandlerTestSuite))
}

func (suite *ListVehicleHandlerTestSuite) SetupSubTest() {
	setup()
	router.HandleFunc("/vehicles", handler.HandleListVehicles).Methods(http.MethodGet)
}

func (suite *ListVehicleHandlerTestSuite) TearDownSubTest() {
	listVehicles.AssertExpectations(suite.T())
}

func (suite *ListVehicleHandlerTestSuite) TestHandleListVehicles() {
	suite.Run("Success", func() {
		expectedVehicles := []domains.Vehicle{
			{ID: "1", Brand: "Toyota", Model: "Corolla", Year: 2020, Color: "Blue", Price: 20000},
		}
		listVehicles.On("Execute", mock.Anything).Return(expectedVehicles, nil)

		req := httptest.NewRequest(http.MethodGet, "/vehicles", nil)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		suite.Equal(http.StatusOK, rr.Code)
		var response []responses.Vehicle
		err := json.NewDecoder(rr.Body).Decode(&response)
		suite.NoError(err)
		suite.Equal(len(expectedVehicles), len(response))
		suite.Equal(expectedVehicles[0].ID, response[0].Id)
		suite.Equal(expectedVehicles[0].Brand, response[0].Brand)
		suite.Equal(expectedVehicles[0].Model, response[0].Model)
		suite.Equal(expectedVehicles[0].Year, response[0].Year)
		suite.Equal(expectedVehicles[0].Color, response[0].Color)
		suite.Equal(expectedVehicles[0].Price, response[0].Price)

	})

	suite.Run("Should return error", func() {
		listVehicles.On("Execute", mock.Anything).Return([]domains.Vehicle{}, errors.New("some error"))

		req := httptest.NewRequest(http.MethodGet, "/vehicles", nil)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		suite.Equal(http.StatusInternalServerError, rr.Code)
		var problemDetails http_errors.ProblemDetails
		err := json.NewDecoder(rr.Body).Decode(&problemDetails)
		suite.NoError(err)
		suite.Contains(problemDetails.Detail, "Failed to list vehicles")
	})
}
