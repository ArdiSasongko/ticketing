package admin_service

import (
	"github.com/ArdiSasongko/ticketing_app/helper"
	admin_entity "github.com/ArdiSasongko/ticketing_app/model/entity/admin"
	admin_web "github.com/ArdiSasongko/ticketing_app/model/web/admin"
)

type AuthService interface {
	Register(req admin_web.RegisterAdminRequest) (helper.CustomResponse, error)
	Login(req admin_web.LoginAdminRequest) (helper.CustomResponse, error)
	Update(userId int, req admin_web.UpdateAdminRequest) (helper.CustomResponse, error)
	ViewMe(userId int) (admin_entity.AdminEntity, error)
}
