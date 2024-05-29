package service

import (
	entity "github.com/ArdiSasongko/ticketing_app/model/entity/event"
	repository "github.com/ArdiSasongko/ticketing_app/repository/event"
)

type EventServiceImpl struct {
	repository   repository.EventRepository
	tokenUseCase helper.TokenUseCase
}

func (service *EventServiceImpl) GetEventList() ([]entity.EventEntity, error) {
	getEventList, errGetEventList := service.repository.GetEvents()

	if errGetEventList != nil {
		return []entity.EventEntity{}, errGetEventList
	}

	return entity.ToEventListEntity(getEventList), nil
}
