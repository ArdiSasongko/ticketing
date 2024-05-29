package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type EventControllerImpl struct {
	EventService service.EventService
}

func (controller *EventControllerImpl) GetEventList(c echo.Context) error {
	getEvents, errGetEvents := controller.EventService.GetEventList()

	if errGetEvents != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, errGetusers.Error(), nil))
	}
	fmt.Println(getEvents)
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "success", getEvents))
}

func (controller *EventControllerImpl) DeleteEvent(c echo.Context) error {

	event := new(web.EventDeleteServiceRequest)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(event); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	deleteEvent, errDeleteEvent := controller.EventService.DeleteEvent(*event, id)

	if errDeleteEvent != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errDeleteEvent.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Event berhasil dihapus", deleteEvent))
}
