package main

import (
	"context"
	"errors"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"log"
	"net/http"
	"vehicle-registration-manager/cmd/vehicle-registration-manager/configs"
	routes "vehicle-registration-manager/internal/adapters/http"
	vehicleHttp "vehicle-registration-manager/internal/adapters/http/handlers"
	"vehicle-registration-manager/internal/adapters/repository"
	"vehicle-registration-manager/internal/app/usecase"
)

//	@title			Vehicle Registration Manager API
//	@version		1.0
//	@description	This is a sample server Vehicle server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Vehicle API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8080
// @BasePath	/
func main() {
	app := fx.New(
		fx.Provide(
			repository.NewVehicleRepository,
			usecase.NewRegisterVehicle,
			usecase.NewUpdateVehicle,
			usecase.NewListVehicles,
			usecase.NewSearchVehicle,
			vehicleHttp.NewVehicleHandler,
			mux.NewRouter,
		),
		fx.Invoke(
			configs.RegisterHealthCheckRoutes,
			configs.RegisterSwaggerRoutes,
			routes.RegisterRoutes),

		fx.Invoke(registerHooks),
	)

	app.Run()
}

func registerHooks(lifecycle fx.Lifecycle, router *mux.Router) {
	var srv *http.Server
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			srv = &http.Server{
				Addr:    ":8080",
				Handler: router,
			}

			go func() {
				if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Printf("HTTP server ListenAndServe: %v", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
}
