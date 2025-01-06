package requests

import (
	"vehicle-registration-manager/internal/core/domains/consts"
)

type Vehicle struct {
	Brand        string        `json:"brand"`
	Model        string        `json:"model"`
	Color        string        `json:"color"`
	LicensePlate string        `json:"license_plate"`
	Year         int           `json:"year"`
	Price        float64       `json:"price"`
	Status       consts.Status `json:"status"`
}
