package seller_web

type SellerLoginRequest struct {
	Email    string `validate:"email" json:"email"`
	Password string `validate:"required" json:"password"`
}

type SellerServiceRequest struct {
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}
