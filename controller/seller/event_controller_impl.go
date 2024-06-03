package seller_controller

import (
	"net/http"

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
