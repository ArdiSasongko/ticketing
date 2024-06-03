package buyer_service

import "github.com/ArdiSasongko/ticketing_app/model/entity/buyer"

type EventService interface {
	GetEventList(filters map[string]string, sort string, limit int, page int) ([]buyer_entity.EventEntity, error)
}
