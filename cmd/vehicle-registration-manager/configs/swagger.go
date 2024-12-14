package configs

import (
	_ "github.com/ArianeASA/vehicle-registration-manager/docs" // This line is necessary for go-swagger to find your docs!
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func RegisterSwaggerRoutes(router *mux.Router) {
	//router.PathPrefix("/swagger/").
	//	Handler(http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swagger")))).
	//	Methods("GET")

	//router.Handle("/swagger/*", httpSwagger.Handler(
	//	httpSwagger.URL("http://localhost:8080/swagger/docs/swagger.json"), //The url pointing to API definition
	//)).Methods("GET")
	//router.HandleFunc("/swagger/*", httpSwagger.WrapHandler.ServeHTTP).Methods("GET")

	//router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
	//	httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
	//	httpSwagger.DeepLinking(true),
	//	httpSwagger.DocExpansion("none"),
	//	httpSwagger.DomID("swagger-ui"),
	//)).Methods(http.MethodGet)
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
}
