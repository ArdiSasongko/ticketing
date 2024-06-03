package seller

import (
	entity "github.com/ArdiSasongko/ticketing_app/model/entity/seller"
	"github.com/ArdiSasongko/ticketing_app/model/web/seller"
)

type SellerService interface {
	SaveSeller(request SellerServiceRequest) (map[string]interface{}, error)
	LoginSeller(email string, password string) (map[string]interface{}, error)
	GetSeller(SellerID int) (entity.SellerEntity, error)
	UpdateSeller(request seller.SellerUpdateServiceRequest, pathId int) (map[string]interface{}, error)
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
