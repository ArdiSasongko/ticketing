package eventcontroller

import (
	"net/http"

	"github.com/ArdiSasongko/ticketing_app/model"
	"github.com/ArdiSasongko/ticketing_app/model/web"
	eventservice "github.com/ArdiSasongko/ticketing_app/pkg/event/event.service"
	"github.com/labstack/echo/v4"
	"strconv"
)

type EventsControllerImpl struct {
	service eventservice.EventsService
}

func NewEventsController(service eventservice.EventsService) *EventsControllerImpl {
	return &EventsControllerImpl{
		service: service,
	}
}

func (controller *EventsControllerImpl) SaveEvents(c echo.Context) error {
	events := new(web.CreateEventsRequest)

	if err := c.Bind(events); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(events); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	saveEvents, errSaveEvents := controller.service.SaveEvents(*events)

	if errSaveEvents != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errSaveEvents.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Successfully created event", saveEvents))
}

//update
func (controller *EventsControllerImpl) UpdateEvent(c echo.Context) error {
    // Mengambil ID acara dari parameter rute
    eventID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, "Invalid event ID", nil))
    }

    // Mengikat data dari permintaan ke struktur UserUpdateServiceRequest
    updateRequest := new(web.UserUpdateServiceRequest)
    if err := c.Bind(updateRequest); err != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
    }

    // Melakukan validasi data yang diikat
    if err := c.Validate(updateRequest); err != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
    }

    // Memanggil layanan untuk melakukan pembaruan acara
    updatedEvent, err := controller.service.UpdateEvent(*updateRequest, eventID)
    if err != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
    }

    // Mengembalikan respons dengan data acara yang diperbarui
    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Successfully updated event", updatedEvent))
}
