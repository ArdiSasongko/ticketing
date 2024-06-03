package event_controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ArdiSasongko/ticketing_app/helper"
	web "github.com/ArdiSasongko/ticketing_app/model/web/event"
	service "github.com/ArdiSasongko/ticketing_app/service/event"
	"github.com/labstack/echo/v4"
)

type EventControllerImpl struct {
	EventService service.EventService
}

func (controller *EventControllerImpl) GetEventList(c echo.Context) error {
	getEvents, errGetEvents := controller.EventService.GetEventList()

	if errGetEvents != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseClient(http.StatusInternalServerError, errGetEvents.Error(), nil))
	}
	fmt.Println(getEvents)
	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "success", getEvents))
}

func (controller *EventControllerImpl) DeleteEvent(c echo.Context) error {

	event := new(web.EventDeleteServiceRequest)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(event); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	deleteEvent, errDeleteEvent := controller.EventService.DeleteEvent(*event, id)

	if errDeleteEvent != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, errDeleteEvent.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Event berhasil dihapus", deleteEvent))
}
