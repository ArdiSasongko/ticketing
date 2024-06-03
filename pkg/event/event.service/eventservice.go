package eventservice

import (
	"github.com/ArdiSasongko/ticketing_app/model/web"
	"github.com/ArdiSasongko/ticketing_app/model/domain"
)

type EventsService interface {
	SaveEvents(request web.CreateEventsRequest) (map[string]interface{}, error)
	GetEventByID(id int) (domain.Event, error) // Tambahkan metode ini
	UpdateEvent(request   web.UserUpdateServiceRequest, pathId int) (map[string]interface{}, error)
	
}
