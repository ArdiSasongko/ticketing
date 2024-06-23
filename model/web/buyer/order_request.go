package buyer_web

type CreateOrderRequest struct {
	EventID int `validate:"required,numeric" json:"event_id"`
	Qty     int `validate:"required,numeric" json:"qty"`
}
