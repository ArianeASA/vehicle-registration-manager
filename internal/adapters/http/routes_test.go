package http_test

import (
	"github.com/stretchr/testify/suite"
	"testing"
	httpAdapt "vehicle-registration-manager/internal/adapters/http"
)

type RegisterRoutesTestSuite struct {
	suite.Suite
}

func TestRegisterRoutesTestSuite(t *testing.T) {
	suite.Run(t, new(RegisterRoutesTestSuite))
}

func (suite *RegisterRoutesTestSuite) SetupSubTest() {
	setup()
	httpAdapt.RegisterRoutes(router, handler)
}

func (suite *RegisterRoutesTestSuite) TestRegisterRoutes() {
	suite.Run("Should register all vehicles route", func() {
		listVehicles := router.GetRoute("ListVehicles")
		suite.NotNil(listVehicles)
		searchVehicleByID := router.GetRoute("SearchVehicleByID")
		suite.NotNil(searchVehicleByID)
		updateVehicle := router.GetRoute("UpdateVehicle")
		suite.NotNil(updateVehicle)
		createVehicle := router.GetRoute("CreateVehicle")
		suite.NotNil(createVehicle)
	})
}
