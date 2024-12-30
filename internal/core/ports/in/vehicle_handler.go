package in

import (
	"github.com/stretchr/testify/mock"
	"net/http"
)

type VehicleHandler interface {
	HandleCreateVehicle(w http.ResponseWriter, r *http.Request)
	HandleUpdateVehicle(w http.ResponseWriter, r *http.Request)
	HandleListVehicles(w http.ResponseWriter, r *http.Request)
	HandleSearchVehicleByID(w http.ResponseWriter, r *http.Request)
}

// MockVehicleHandler is a mock of VehicleHandler
type MockVehicleHandler struct {
	mock.Mock
}

func (m *MockVehicleHandler) HandleCreateVehicle(w http.ResponseWriter, r *http.Request) {
	m.Called(w, r)
}

func (m *MockVehicleHandler) HandleUpdateVehicle(w http.ResponseWriter, r *http.Request) {
	m.Called(w, r)
}

func (m *MockVehicleHandler) HandleListVehicles(w http.ResponseWriter, r *http.Request) {
	m.Called(w, r)
}

func (m *MockVehicleHandler) HandleSearchVehicleByID(w http.ResponseWriter, r *http.Request) {
	m.Called(w, r)
}
