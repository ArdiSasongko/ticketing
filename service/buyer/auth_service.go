package buyer_service

import (
	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/model/web/buyer"
)

type BuyerServiceInterface interface {
	Register(req buyer_web.BuyerRequest) (helper.CustomResponse, error)
	Login(email, password string) (helper.CustomResponse, error)
	Update(userId int, req buyer_web.BuyerUpdateRequest) (helper.CustomResponse, error)
}
