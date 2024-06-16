package seller_service

import (
	"errors"
	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/model/domain"
	"github.com/ArdiSasongko/ticketing_app/model/entity/buyer"
	"github.com/ArdiSasongko/ticketing_app/model/entity/seller"
	"github.com/ArdiSasongko/ticketing_app/model/enum"
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

func (service *EventServiceImpl) ViewEvent(eventId int) (buyer_entity.EventEntity, error) {
	event, err := service.repository.GetEventByID(eventId)
	if err != nil {
		return buyer_entity.EventEntity{}, err
	}

	return buyer_entity.ToEventEntity(event), nil
}

func (service *EventServiceImpl) SaveEvents(userID int, request seller_web.CreateEventsRequest) (map[string]interface{}, error) {
	date, err := helper.ParseDate(request.Date)
	if err != nil {
		return nil, err
	}

	if err := helper.ValidateName(request.Name); err != nil {
		return nil, err
	}

	if err := helper.ValidateCategory(request.Category); err != nil {
		return nil, err
	}
	if err := helper.ValidateLocation(request.Location); err != nil {
		return nil, err
	}
	if err := helper.ValidateQty(request.Qty); err != nil {
		return nil, err
	}
	if err := helper.ValidatePrice(request.Price); err != nil {
		return nil, err
	}

	eventReq := domain.Event{
		SellerID: userID,
		Name:     request.Name,
		Date:     date,
		Location: request.Location,
		Qty:      request.Qty,
		Category: request.Category,
		Price:    request.Price,
		Status:   enum.EventStatusInactive,
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
		"status":    saveEvent.Status,
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

func (service *EventServiceImpl) UpdateEventStatus(request seller_web.UpdateEventStatusRequest, id int) (seller_entity.EventEntity, error) {
	event, getEventErr := service.repository.GetEventByID(id)
	if getEventErr != nil {
		return seller_entity.EventEntity{}, getEventErr
	}

	if event.Status == enum.EventStatusClosed {
		return seller_entity.EventEntity{}, errors.New("cannot update closed event")
	}

	event.Status = request.Status
	event, updateEventErr := service.repository.UpdateEvent(event)
	if updateEventErr != nil {
		return seller_entity.EventEntity{}, updateEventErr
	}

	return seller_entity.ToEventEntity(event), nil
}

func (service *EventServiceImpl) GetEventByID(eventID int) (domain.Event, error) {
	return service.repository.GetEventByID(eventID)
}

func (service *EventServiceImpl) CheckInTicket(eventID int, ticketID int) error {
	err := service.repository.CheckInTicket(eventID, ticketID)
	if err != nil {
		return err
	}
	return nil
}

func (service *EventServiceImpl) DeleteEvent(eventId int) error {
	return service.repository.DeleteEventById(eventId)
}
