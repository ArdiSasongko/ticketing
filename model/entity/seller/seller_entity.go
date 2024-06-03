package seller_entity

type SellerEntity struct {
	SellerID int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func ToSellerEntity(id int, name string, email string) SellerEntity {
	return SellerEntity{
		SellerID: id,
		Name:     name,
		Email:    email,
	}
}
