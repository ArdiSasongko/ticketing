package seller_service

import (
	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"github.com/ArdiSasongko/ticketing_app/model/entity/buyer"
	"github.com/ArdiSasongko/ticketing_app/model/web/seller"
	"github.com/ArdiSasongko/ticketing_app/repository/seller"
)

type EventServiceImpl struct {
	repository seller_repository.EventRepository
}

func NewEventService(repository seller_repository.EventRepository) *EventServiceImpl {
	return &EventServiceImpl{
		repository: repository,
	}
}

func (service *EventServiceImpl) GetEventList(sellerId int, filters map[string]string, sort string, limit int, page int) ([]buyer_entity.EventEntity, error) { // todo
	events, err := service.repository.ListEvents(sellerId, filters, sort, limit, page)
	if err != nil {
		return []buyer_entity.EventEntity{}, err
	}

	return buyer_entity.ToEventListEntity(events), nil
}

func (service *EventServiceImpl) SaveEvents(request seller_web.CreateEventsRequest) (map[string]interface{}, error) {
	date, err := helper.ParseDate(request.Date)
	if err != nil {
		return nil, err
	}

	eventReq := domain.Event{
		SellerID: request.SellerID,
		Name:     request.Name,
		Date:     date,
		Location: request.Location,
		Qty:      request.Qty,
		Category: request.Category,
		Price:    request.Price,
	}

	saveEvent, errSaveEvent := service.repository.CreateEvent(eventReq)

	if errSaveEvent != nil {
		return nil, errSaveEvent
	}

	return map[string]interface{}{
		"seller_id": saveEvent.SellerID,
		"name":      saveEvent.Name,
		"date":      saveEvent.Date,
		"location":  saveEvent.Location,
		"qty":       saveEvent.Qty,
		"category":  saveEvent.Category,
		"price":     saveEvent.Price,
	}, nil
}

func (service *EventServiceImpl) UpdateEvent(request seller_web.UserUpdateServiceRequest, eventID int) (map[string]interface{}, error) {
	// Mengambil data acara yang akan diperbarui berdasarkan ID
	existingEvent, err := service.repository.GetEventByID(eventID)
	if err != nil {
		return nil, err
	}

	// Memeriksa setiap bidang dalam permintaan pembaruan.
	// Jika bidang tidak diisi dalam permintaan, maka tetap menggunakan nilai yang sudah ada.
	if request.Name == "" {
		request.Name = existingEvent.Name
	}
	if request.Date == "" {
		request.Date = existingEvent.Date.Format("2006-01-02") // Konversi ke string
	}
	if request.Location == "" {
		request.Location = existingEvent.Location
	}
	if request.Qty == 0 {
		request.Qty = existingEvent.Qty
	}
	if request.Category == "" {
		request.Category = existingEvent.Category
	}
	if request.Price == 0 {
		request.Price = existingEvent.Price
	}

	// Membuat objek domain.Event baru berdasarkan data yang diperbarui
	date, err := helper.ParseDate(request.Date)
	if err != nil {
		return nil, err
	}

	updatedEvent := domain.Event{
		ID:       existingEvent.ID,
		SellerID: existingEvent.SellerID,
		Name:     request.Name,
		Date:     date,
		Location: request.Location,
		Qty:      request.Qty,
		Category: request.Category,
		Price:    request.Price,
	}

	// Memanggil repository untuk melakukan pembaruan acara
	updatedEvent, err = service.repository.UpdateEvent(updatedEvent)
	if err != nil {
		return nil, err
	}

	// Mengembalikan respons dengan data acara yang diperbarui dalam bentuk map
	return map[string]interface{}{
		"id":        updatedEvent.ID,
		"seller_id": updatedEvent.SellerID,
		"name":      updatedEvent.Name,
		"date":      updatedEvent.Date,
		"location":  updatedEvent.Location,
		"qty":       updatedEvent.Qty,
		"category":  updatedEvent.Category,
		"price":     updatedEvent.Price,
	}, nil
}

func (service *EventServiceImpl) GetEventByID(eventID int) (domain.Event, error) {
	return service.repository.GetEventByID(eventID)
}
