package configs

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	_ "vehicle-registration-manager/docs" // This line is necessary for go-swagger to find your docs!
)

func RegisterSwaggerRoutes(router *mux.Router) {
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
}
