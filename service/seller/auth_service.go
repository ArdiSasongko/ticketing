package seller

type SellerService interface {
	SaveSeller(request SellerServiceRequest) (map[string]interface{}, error)
	LoginSeller(email string, password string) (map[string]interface{}, error)
}

type SellerServiceRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SellerLoginRequest struct {
	Email    string `validate:"email" json:"email"`
	Password string `validate:"required" json:"password"`
}
