package buyer_web

type Order struct {
	Qty int `validate:"required" json:"column:qty"`
}
