package entities

type Vehicle struct {
	ID    string  `db:"id"`
	Brand string  `db:"brand"`
	Model string  `db:"model"`
	Color string  `db:"color"`
	Year  int     `db:"year"`
	Price float64 `db:"price"`
}
