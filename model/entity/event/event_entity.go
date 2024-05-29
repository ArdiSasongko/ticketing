type EventEntity struct {
	EventID    int     `json:"event_id"`
	SellerIDFK int     `json:"seller_id_fk`
	Name       string  `json:"name"`
	Date       string  `json:"date"`
	Location   string  `json:"location"`
	Qty        int     `json:"qty"`
	Category   string  `json:"category"`
	Price      float32 `json:"price"`
}