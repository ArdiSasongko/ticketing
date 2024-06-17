package admin_web

type AdminRequest struct {
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}

type AdminLoginRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}
