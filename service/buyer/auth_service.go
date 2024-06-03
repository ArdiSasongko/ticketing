package buyer

import (
	"github.com/ArdiSasongko/ticketing_app/helper"
	buyerreq "github.com/ArdiSasongko/ticketing_app/model/web/buyer"
)

type BuyerServiceInterface interface {
	Register(req buyerreq.BuyerRequest) (helper.CustomResponse, error)
	Login(email, password string) (helper.CustomResponse, error)
	Update(userId int, req buyerreq.BuyerUpdateRequest) (helper.CustomResponse, error)
}
