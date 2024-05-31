package service

import (
	"github.com/ArdiSasongko/ticketing_app/helper"
	entity "github.com/ArdiSasongko/ticketing_app/model/entity/event"
	repository "github.com/ArdiSasongko/ticketing_app/repository/event"
)

type EventServiceImpl struct {
	repository     repository.EventRepository
	CustomResponse helper.CustomResponse
}

func (service *EventServiceImpl) GetEventList() ([]entity.EventEntity, error) {
	getEventList, errGetEventList := service.repository.GetEvents()

	if errGetEventList != nil {
		return []entity.EventEntity{}, errGetEventList
	}

	return entity.ToEventListEntity(getEventList), nil
}
