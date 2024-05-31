package seller

import web "github.com/ArdiSasongko/ticketing_app/model/web/seller"

type SellerService interface {
	UpdateSeller(request web.SellerUpdateServiceRequest, pathId int) (map[string]interface{}, error)
}
