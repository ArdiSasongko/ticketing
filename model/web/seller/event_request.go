package seller_web

type CreateEventRequest struct {
	Name     string  `json:"name" validate:"required"`
	Date     string  `json:"date" validate:"required"`
	Location string  `json:"location" validate:"required"`
	Qty      int     `json:"qty" validate:"required"`
	Category string  `json:"category" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
}

type UpdateEventRequest struct {
	Name     string  `json:"name" validate:"required"`
	Date     string  `json:"date" validate:"required"`
	Location string  `json:"location" validate:"required"`
	Qty      int     `json:"qty" validate:"required"`
	Category string  `json:"category" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
}

type UpdateEventStatusRequest struct {
	Status string `json:"status" validate:"required,event_status_enum"`
}
