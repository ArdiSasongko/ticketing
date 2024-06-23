package buyer_service

import (
	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/model/entity/buyer"
	"github.com/ArdiSasongko/ticketing_app/model/web/buyer"
)

type AuthService interface {
	Register(req buyer_web.RegisterBuyerRequest) (helper.CustomResponse, error)
	Login(email, password string) (helper.CustomResponse, error)
	Update(userId int, req buyer_web.UpdateBuyerRequest) (helper.CustomResponse, error)
	ViewMe(userId int) (buyer_entity.BuyerEntity, error)
}
