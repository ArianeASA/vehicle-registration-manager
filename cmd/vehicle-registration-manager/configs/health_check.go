package configs

import (
	"github.com/gorilla/mux"
	"net/http"
	configsDB "vehicle-registration-manager/internal/adapters/repository/configs"
	httpErrors "vehicle-registration-manager/pkg/http_errors"
	"vehicle-registration-manager/pkg/tracer"
)

var dbConfig *configsDB.DatabaseConfig

func RegisterHealthCheckRoutes(router *mux.Router, config *configsDB.DatabaseConfig) {
	dbConfig = config
	router.HandleFunc("/health", HealthCheck).Methods(http.MethodGet)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	trc := tracer.NewTracer(r)
	if err := dbConfig.Ping(); err != nil {
		msg := "Database is not available"
		trc.Logger.Error(msg, err)
		httpErrors.WriteProblemDetails(w, httpErrors.InternalServerError(msg, r.RequestURI))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Service is healthy!"))
}
