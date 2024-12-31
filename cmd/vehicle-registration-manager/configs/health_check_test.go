package configs_test

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
	configsMain "vehicle-registration-manager/cmd/vehicle-registration-manager/configs"
	"vehicle-registration-manager/internal/adapters/repository/configs"
)

type HealthCheckTestSuite struct {
	suite.Suite
	router *mux.Router
	config *configs.MockDatabaseConfig
}

func TestHealthCheckTestSuite(t *testing.T) {
	suite.Run(t, new(HealthCheckTestSuite))
}

func (suite *HealthCheckTestSuite) SetupTest() {
	suite.router = mux.NewRouter()
	suite.config = new(configs.MockDatabaseConfig)

	configsMain.RegisterHealthCheckRoutes(suite.router, suite.config)
}

func (suite *HealthCheckTestSuite) TearDownTest() {
	suite.config.AssertExpectations(suite.T())
}

const uri = "/health"

func (suite *HealthCheckTestSuite) TestHealthCheckSuccess() {
	suite.config.On("Ping").Return(nil)
	req := httptest.NewRequest(http.MethodGet, uri, nil)
	rr := httptest.NewRecorder()

	suite.router.ServeHTTP(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)
	assert.Equal(suite.T(), "Service is healthy!", rr.Body.String())
}

func (suite *HealthCheckTestSuite) TestHealthCheckFailure() {
	suite.config.On("Ping").Return(errors.New("database error"))
	req := httptest.NewRequest(http.MethodGet, uri, nil)
	rr := httptest.NewRecorder()

	suite.router.ServeHTTP(rr, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "Database is not available")
}
