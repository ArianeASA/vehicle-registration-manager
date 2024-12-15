package domain

import "strings"

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
