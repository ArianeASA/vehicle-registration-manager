package domains

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"vehicle-registration-manager/internal/core/domains/consts"
)

var (
	ErrVehicleNotFound = errors.New("vehicle not found")
	validNextStatus    = map[consts.Status][]consts.Status{
		consts.StatusForSale:              {consts.StatusReserved, consts.StatusCanceled},
		consts.StatusReserved:             {consts.StatusSold, consts.StatusCancelingReservation},
		consts.StatusCancelingReservation: {consts.StatusForSale},
	}
)

const (
	stringEmpty = ""
)

type Vehicle struct {
	ID           string
	Brand        string
	Model        string
	Color        string
	LicensePlate string
	Year         int
	Price        float64
	Status       consts.Status
}

func (v *Vehicle) Exist() bool {
	return !strings.EqualFold(v.ID, stringEmpty)
}

func (v *Vehicle) IsValidCreate() bool {
	return !strings.EqualFold(v.LicensePlate, stringEmpty) &&
		!strings.EqualFold(v.Brand, stringEmpty) &&
		!strings.EqualFold(v.Model, stringEmpty) &&
		!strings.EqualFold(v.Color, stringEmpty) &&
		v.Year != 0 &&
		v.Price != 0 &&
		!strings.EqualFold(string(v.Status), stringEmpty)
}

func (v *Vehicle) isValidUpdateStatus(newStatus consts.Status) bool {
	return slices.Contains(validNextStatus[v.Status], newStatus)
}

func (v *Vehicle) UpdateFields(vehicle *Vehicle) error {
	if !strings.EqualFold(stringEmpty, vehicle.Brand) && !strings.EqualFold(v.Brand, vehicle.Brand) {
		v.Brand = vehicle.Brand
	}

	if !strings.EqualFold(stringEmpty, vehicle.Model) && !strings.EqualFold(v.Model, vehicle.Model) {
		v.Model = vehicle.Model
	}

	if vehicle.Year != 0 {
		v.Year = vehicle.Year
	}

	if !strings.EqualFold(stringEmpty, vehicle.Color) && !strings.EqualFold(v.Color, vehicle.Color) {
		v.Color = vehicle.Color
	}

	if vehicle.Price != 0 {
		v.Price = vehicle.Price
	}

	if !strings.EqualFold(stringEmpty, string(vehicle.Status)) && v.Status != vehicle.Status {
		if !v.isValidUpdateStatus(vehicle.Status) {
			return fmt.Errorf("invalid status transition from %s to %s. Rule status permitted: from %s to %v",
				v.Status, vehicle.Status, v.Status, validNextStatus[v.Status])
		}
		v.Status = vehicle.Status
	}

	return nil
}
