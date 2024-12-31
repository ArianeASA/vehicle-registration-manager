package in

import (
	"github.com/stretchr/testify/mock"
	"net/http"
)

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
