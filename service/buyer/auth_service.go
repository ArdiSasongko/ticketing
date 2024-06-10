package buyer_service

import (
	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/model/entity/buyer"
	"github.com/ArdiSasongko/ticketing_app/model/entity/history"
	"github.com/ArdiSasongko/ticketing_app/model/web/buyer"
)

type BuyerServiceInterface interface {
	Register(req buyer_web.BuyerRequest) (helper.CustomResponse, error)
	Login(email, password string) (helper.CustomResponse, error)
	Update(userId int, req buyer_web.BuyerUpdateRequest) (helper.CustomResponse, error)
	GetAll() ([]buyer_entity.BuyerEntity, error)
	GetHistory(userId int) ([]history_entity.HistoryEntity, error)
}
