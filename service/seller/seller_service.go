package seller

import (
	entity "github.com/ArdiSasongko/ticketing_app/model/entity/seller"
	web "github.com/ArdiSasongko/ticketing_app/model/web/seller"
)

type SellerService interface {
	GetSeller(SellerID int) (entity.SellerEntity, error)
	UpdateSeller(request web.SellerUpdateServiceRequest, pathId int) (map[string]interface{}, error)
}
