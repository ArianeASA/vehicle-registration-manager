package entities

type Vehicle struct {
	ID           string  `db:"id"`
	Brand        string  `db:"brand"`
	Model        string  `db:"model"`
	Color        string  `db:"color"`
	Status       string  `db:"status"`
	LicensePlate string  `db:"license_plate"`
	Year         int     `db:"year"`
	Price        float64 `db:"price"`
}
