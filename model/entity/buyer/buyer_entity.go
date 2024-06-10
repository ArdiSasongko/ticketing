package buyer_entity

import "time"

type BuyerEntity struct {
	BuyerID   int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
