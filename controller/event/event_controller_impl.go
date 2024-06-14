package event_controller

import (
	"github.com/ArdiSasongko/ticketing_app/helper"
	service "github.com/ArdiSasongko/ticketing_app/service/event"
	"github.com/labstack/echo/v4"
	"net/http"
)

type EventControllerImpl struct {
	EventService service.EventService
}

func (controller *EventControllerImpl) GetEventList(c echo.Context) error {
	getEvents, errGetEvents := controller.EventService.GetEventList()

	if errGetEvents != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseClient(http.StatusInternalServerError, errGetEvents.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "success", getEvents))
}
