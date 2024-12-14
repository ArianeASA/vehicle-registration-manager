package requests

type Vehicle struct {
	Brand string  `json:"brand"`
	Model string  `json:"model"`
	Year  int     `json:"year"`
	Color string  `json:"color"`
	Price float64 `json:"price"`
}
