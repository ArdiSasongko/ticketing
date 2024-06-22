package admin_service

import (
	"github.com/ArdiSasongko/ticketing_app/helper"
	buyer_entity "github.com/ArdiSasongko/ticketing_app/model/entity/buyer"
	seller_entity "github.com/ArdiSasongko/ticketing_app/model/entity/seller"
	admin_web "github.com/ArdiSasongko/ticketing_app/model/web/admin"
)

type AdminServiceInterface interface {
	Register(req admin_web.AdminRequest) (helper.CustomResponse, error)
	Login(email, password string) (helper.CustomResponse, error)
	GetBuyers() ([]buyer_entity.BuyerEntity, error)
	GetSellers() ([]seller_entity.SellerEntity, error)
	GetBuyerByID(buyerID int) (buyer_entity.BuyerEntity, error)
	GetSellerByID(sellerID int) (seller_entity.SellerEntity, error)
}
