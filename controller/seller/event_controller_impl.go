package seller_controller

import (
	"github.com/ArdiSasongko/ticketing_app/model"
	"github.com/ArdiSasongko/ticketing_app/model/web/seller"
	"net/http"
	"strconv"

	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/service/seller"
	"github.com/labstack/echo/v4"
)

type EventControllerImpl struct {
	eventService seller_service.EventService
}

func NewEventController(service seller_service.EventService) *EventControllerImpl {
	return &EventControllerImpl{
		eventService: service,
	}
}

func (controller *EventControllerImpl) GetEventList(c echo.Context) error {
	filters, sort, limit, page := helper.ExtractFilterSort(c.QueryParams())
	sellerId := 1 // TODO: auth
	events, err := controller.eventService.GetEventList(sellerId, filters, sort, limit, page)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseClient(http.StatusInternalServerError, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "success", events))
}

func (controller *EventControllerImpl) SaveEvents(c echo.Context) error {
	events := new(seller_web.CreateEventsRequest)

	if err := c.Bind(events); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}


	saveEvents, errSaveEvents := controller.eventService.SaveEvents(*events)

	if errSaveEvents != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errSaveEvents.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Successfully created event", saveEvents))
}

// update
func (controller *EventControllerImpl) UpdateEvent(c echo.Context) error {
	// Mengambil ID acara dari parameter rute
	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, "Invalid event ID", nil))
	}

	// Mengikat data dari permintaan ke struktur UserUpdateServiceRequest
	updateRequest := new(seller_web.UserUpdateServiceRequest)
	if err := c.Bind(updateRequest); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}



	// Memanggil layanan untuk melakukan pembaruan acara
	updatedEvent, err := controller.eventService.UpdateEvent(*updateRequest, eventID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	// Mengembalikan respons dengan data acara yang diperbarui
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Successfully updated event", updatedEvent))
}
