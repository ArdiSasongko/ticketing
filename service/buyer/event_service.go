package buyer

import entity "github.com/ArdiSasongko/ticketing_app/model/entity/buyer"

type EventService interface {
	GetEventList(filters map[string]string, sort string, limit int, page int) ([]entity.EventEntity, error)
}
