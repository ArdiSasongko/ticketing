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
	eventRepository seller_repository.EventRepository
}

func NewEventService(eventRepository seller_repository.EventRepository) *EventServiceImpl {
	return &EventServiceImpl{
		eventRepository: eventRepository,
	}
}

func (service *EventServiceImpl) GetEventList(sellerId int, filters map[string]string, sort string, limit int, page int) ([]buyer_entity.EventEntity, error) { // todo
	events, err := service.eventRepository.ListEvents(sellerId, filters, sort, limit, page)
	if err != nil {
		return []buyer_entity.EventEntity{}, err
	}

	return buyer_entity.ToEventListEntity(events), nil
}

func (service *EventServiceImpl) ViewEvent(eventId int) (buyer_entity.EventEntity, error) {
	event, err := service.eventRepository.GetEventByID(eventId)
	if err != nil {
		return buyer_entity.EventEntity{}, err
	}

	return buyer_entity.ToEventEntity(event), nil
}

func (service *EventServiceImpl) SaveEvents(userID int, request seller_web.CreateEventRequest) (seller_entity.EventEntity, error) {
	date, err := helper.ParseDate(request.Date)
	if err != nil {
		return seller_entity.EventEntity{}, err
	}

	if err := helper.ValidateName(request.Name); err != nil {
		return seller_entity.EventEntity{}, err
	}
	if err := helper.ValidateCategory(request.Category); err != nil {
		return seller_entity.EventEntity{}, err
	}
	if err := helper.ValidateLocation(request.Location); err != nil {
		return seller_entity.EventEntity{}, err
	}
	if err := helper.ValidateQty(request.Qty); err != nil {
		return seller_entity.EventEntity{}, err
	}
	if err := helper.ValidatePrice(request.Price); err != nil {
		return seller_entity.EventEntity{}, err
	}

	eventReq := domain.Event{
		SellerID: userID,
		Name:     request.Name,
		Date:     date,
		Location: request.Location,
		Qty:      request.Qty,
		Category: request.Category,
		Price:    request.Price,
		Status:   string(enum.EventStatusInactive),
	}

	saveEvent, errSaveEvent := service.eventRepository.CreateEvent(eventReq)
	if errSaveEvent != nil {
		return seller_entity.EventEntity{}, errSaveEvent
	}

	saveEvent, err = service.eventRepository.GetEventByID(saveEvent.EventID)
	if err != nil {
		return seller_entity.EventEntity{}, err
	}

	return seller_entity.ToEventEntity(saveEvent), nil
}

func (service *EventServiceImpl) UpdateEvent(request seller_web.UpdateEventRequest, eventID int) (seller_entity.EventEntity, error) {
	// Mengambil data acara yang akan diperbarui berdasarkan ID
	existingEvent, err := service.eventRepository.GetEventByID(eventID)
	if err != nil {
		return seller_entity.EventEntity{}, err
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
		return seller_entity.EventEntity{}, err
	}

	updatedEvent := domain.Event{
		EventID:  existingEvent.EventID,
		SellerID: existingEvent.SellerID,
		Name:     request.Name,
		Date:     date,
		Location: request.Location,
		Qty:      request.Qty,
		Category: request.Category,
		Price:    request.Price,
	}

	// Memanggil repository untuk melakukan pembaruan acara
	updatedEvent, err = service.eventRepository.UpdateEvent(updatedEvent)
	if err != nil {
		return seller_entity.EventEntity{}, err
	}

	updatedEvent, err = service.eventRepository.GetEventByID(updatedEvent.EventID)
	if err != nil {
		return seller_entity.EventEntity{}, err
	}

	// Mengembalikan respons dengan data acara yang diperbarui dalam bentuk map
	return seller_entity.ToEventEntity(updatedEvent), nil
}

func (service *EventServiceImpl) UpdateEventStatus(request seller_web.UpdateEventStatusRequest, id int) (seller_entity.EventEntity, error) {
	event, getEventErr := service.eventRepository.GetEventByID(id)
	if getEventErr != nil {
		return seller_entity.EventEntity{}, getEventErr
	}

	if event.Status == string(enum.EventStatusClosed) {
		return seller_entity.EventEntity{}, errors.New("cannot update closed event")
	}

	event.Status = request.Status
	event, updateEventErr := service.eventRepository.UpdateEvent(event)
	if updateEventErr != nil {
		return seller_entity.EventEntity{}, updateEventErr
	}

	return seller_entity.ToEventEntity(event), nil
}

func (service *EventServiceImpl) GetEventByID(eventID int) (domain.Event, error) {
	return service.eventRepository.GetEventByID(eventID)
}

func (service *EventServiceImpl) CheckInTicket(eventID int, ticketID int) error {
	err := service.eventRepository.CheckInTicket(eventID, ticketID)
	if err != nil {
		return err
	}
	return nil
}

func (service *EventServiceImpl) DeleteEvent(eventId int) error {
	return service.eventRepository.DeleteEventById(eventId)
}
