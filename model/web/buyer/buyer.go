package buyer

type BuyerRequest struct {
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}

type BuyerLoginRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}

type BuyerUpdateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
