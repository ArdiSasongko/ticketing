package seller_service

import (
	"github.com/ArdiSasongko/ticketing_app/model/entity/seller"
	"github.com/ArdiSasongko/ticketing_app/model/web/seller"
)

type AuthService interface {
	SaveSeller(request seller_web.RegisterSellerRequest) (map[string]interface{}, error)
	LoginSeller(email string, password string) (map[string]interface{}, error)
	GetSeller(SellerID int) (seller_entity.SellerEntity, error)
	UpdateSeller(request seller_web.UpdateSellerRequest, pathId int) (map[string]interface{}, error)
}
