package out

import (
	"vehicle-registration-manager/internal/core/domains"
	"vehicle-registration-manager/pkg/tracer"
)

type VehicleRepository interface {
	Save(tcr *tracer.Tracer, vehicle domains.Vehicle) error
	Update(tcr *tracer.Tracer, vehicle domains.Vehicle) error
	FindAll(tcr *tracer.Tracer) ([]domains.Vehicle, error)
	FindByID(tcr *tracer.Tracer, id string) (domains.Vehicle, error)
}
