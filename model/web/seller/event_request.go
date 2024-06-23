package seller_web

type CreateEventRequest struct {
	SellerID int     `json:"seller_id" validate:"required"`
	Name     string  `json:"name" validate:"required"`
	Date     string  `json:"date" validate:"required"`
	Location string  `json:"location" validate:"required"`
	Qty      int     `json:"qty" validate:"required"`
	Category string  `json:"category" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
}

type UpdateEventRequest struct {
	SellerID int     `json:"seller_id"`
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
