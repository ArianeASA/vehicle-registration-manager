package consts

// Status is a type that represents the status of a vehicle.
type Status string

const (
	StatusSold                 = "SOLD"
	StatusReserved             = "RESERVED"
	StatusCanceled             = "CANCELED"
	StatusForSale              = "FOR_SALE"
	StatusCancelingReservation = "CANCELING_RESERVATION"
)
