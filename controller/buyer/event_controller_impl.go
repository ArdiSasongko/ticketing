package buyer_controller

import (
	"net/http"
	"strconv"

	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/model"
	"github.com/ArdiSasongko/ticketing_app/service/buyer"
	"github.com/labstack/echo/v4"
)

type EventControllerImpl struct {
	eventService buyer_service.EventService
}

func NewEventController(service buyer_service.EventService) *EventControllerImpl {
	return &EventControllerImpl{
		eventService: service,
	}
}

func (controller *EventControllerImpl) GetEventList(c echo.Context) error {
	filters, sort, limit, page := helper.ExtractFilterSort(c.QueryParams())
	events, err := controller.eventService.GetEventList(filters, sort, limit, page)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "success", events))
}

func (controller *EventControllerImpl) ViewEvent(c echo.Context) error {
	eventId, _ := strconv.Atoi(c.Param("id"))

	event, getEventErr := controller.eventService.ViewEvent(eventId) // todo: load seller
	if getEventErr != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, getEventErr.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "success", event))
}
