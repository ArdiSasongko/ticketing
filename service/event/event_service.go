package service

import entity "github.com/ArdiSasongko/ticketing_app/model/entity/event"

type EventService interface {
	GetUserList() ([]entity.EventEntity, error)
	DeleteEvent(request web.EventDeleteServiceRequest, pathId int) (map[string]interface{}, error)
}
