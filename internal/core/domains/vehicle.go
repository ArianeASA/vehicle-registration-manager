package domains

import (
	"errors"
	"strings"
)

var (
	ErrVehicleNotFound = errors.New("vehicle not found")
)

type Vehicle struct {
	ID    string
	Brand string
	Model string
	Year  int
	Color string
	Price float64
}

func (v *Vehicle) Exist() bool {
	return !strings.EqualFold(v.ID, "")
}
