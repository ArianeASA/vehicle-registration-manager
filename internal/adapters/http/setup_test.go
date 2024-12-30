package http_test

import (
	"github.com/gorilla/mux"
	"vehicle-registration-manager/internal/core/ports/in"
)

var (
	handler *in.MockVehicleHandler
	router  *mux.Router
)

func setup() {
	handler = new(in.MockVehicleHandler)
	router = mux.NewRouter()
}
