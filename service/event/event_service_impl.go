package event_service

import (
	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/model/entity/event"
	"github.com/ArdiSasongko/ticketing_app/repository/event"
)

type EventServiceImpl struct {
	repository     event_repository.EventRepository
	CustomResponse helper.CustomResponse
}

func (service *EventServiceImpl) GetEventList() ([]event_entity.EventEntity, error) {
	getEventList, errGetEventList := service.repository.GetEvents()

	if errGetEventList != nil {
		return []event_entity.EventEntity{}, errGetEventList
	}

	return event_entity.ToEventListEntity(getEventList), nil
}
