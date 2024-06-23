package general_controller

import (
	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/service/general"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type EventControllerImpl struct {
	eventService general_service.EventService
}

func NewEventController(eventService general_service.EventService) *EventControllerImpl {
	return &EventControllerImpl{
		eventService: eventService,
	}
}

// List godoc
// @Summary Event (List)
// @Description Event (List)
// @Tags [General] Event
// @Accept json
// @Produce json
// @Param filter[category] query string false "Filter"
// @Param filter[location] query string false "Filter"
// @Param filter[name] query string false "Filter"
// @Param sort query string false "Sort"
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 500 {object} helper.ResponseClientModel
// @Router /general/events [get]
func (controller *EventControllerImpl) List(c echo.Context) error {
	filters, sort, limit, page := helper.ExtractFilterSort(c.QueryParams())
	getEvents, errGetEvents := controller.eventService.GetEventList(filters, sort, limit, page)

	if errGetEvents != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseClient(http.StatusInternalServerError, errGetEvents.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success", getEvents))
}

// View godoc
// @Summary Event (View)
// @Description Event (View)
// @Tags [General] Event
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 500 {object} helper.ResponseClientModel
// @Router /general/events/{id} [get]
func (controller *EventControllerImpl) View(c echo.Context) error {
	eventId, _ := strconv.Atoi(c.Param("id"))

	getEvents, errGetEvents := controller.eventService.GetEvent(eventId)

	if errGetEvents != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseClient(http.StatusInternalServerError, errGetEvents.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success", getEvents))
}
