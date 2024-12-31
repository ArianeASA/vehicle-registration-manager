package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"log"
	"net/http"
	"vehicle-registration-manager/cmd/vehicle-registration-manager/configs"
	routes "vehicle-registration-manager/internal/adapters/http"
	vehicleHttp "vehicle-registration-manager/internal/adapters/http/handlers"
	"vehicle-registration-manager/internal/adapters/repository"
	configsDB "vehicle-registration-manager/internal/adapters/repository/configs"
	"vehicle-registration-manager/internal/app/usecase/create"
	"vehicle-registration-manager/internal/app/usecase/list"
	"vehicle-registration-manager/internal/app/usecase/search"
	usecase "vehicle-registration-manager/internal/app/usecase/update"
	"vehicle-registration-manager/pkg/logger"
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
		//fx.Supply(configsDB.NewDatabaseConfig()),
		fx.Provide(
			configsDB.NewDatabaseConfig,
			repository.VehicleRepositoryFactory,
			create.NewCreateVehicle,
			usecase.NewUpdateVehicle,
			list.NewListVehicles,
			search.NewSearchVehicle,
			vehicleHttp.NewVehicleHandler,
			mux.NewRouter,
		),
		fx.WithLogger(func() fxevent.Logger {
			return logger.NewFxLogger()
		}),
		fx.Invoke(
			func(configsDatabase configsDB.DatabaseConfigs) error {
				if _, err := configsDatabase.InitDatabase(); err != nil {
					return fmt.Errorf("failed to initialize database: %s", err.Error())
				}
				return nil
			},
		),
		fx.Invoke(
			configs.RegisterHealthCheckRoutes,
			configs.RegisterSwaggerRoutes,
			routes.RegisterRoutes,
			registerHooks,
		),
	)
	app.Run()
}

func registerHooks(lifecycle fx.Lifecycle, router *mux.Router, config configsDB.DatabaseConfigs) {
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
			config.Close()
			return srv.Shutdown(ctx)
		},
	})
}
