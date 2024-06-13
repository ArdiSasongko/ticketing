package buyer_web

type OrderRequest struct {
	EventID int `validate:"required" json:"event_id"`
	Qty     int `validate:"required" json:"qty"`
}
