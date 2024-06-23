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

// List godoc
// @Summary Event (List)
// @Description Event (List)
// @Tags [Buyer] Event
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param filter[category] query string false "Filter"
// @Param filter[location] query string false "Filter"
// @Param filter[name] query string false "Filter"
// @Param sort query string false "Sort"
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 500 {object} helper.ResponseClientModel
// @Router /buyer/events [get]
func (controller *EventControllerImpl) List(c echo.Context) error {
	filters, sort, limit, page := helper.ExtractFilterSort(c.QueryParams())
	events, err := controller.eventService.GetEventList(filters, sort, limit, page)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "success", events))
}

// View godoc
// @Summary Event (View)
// @Description Event (View)
// @Tags [Buyer] Event
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param id path int true "id"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 500 {object} helper.ResponseClientModel
// @Router /buyer/events/{id} [get]
func (controller *EventControllerImpl) View(c echo.Context) error {
	eventId, _ := strconv.Atoi(c.Param("id"))

	event, getEventErr := controller.eventService.ViewEvent(eventId)
	if getEventErr != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, getEventErr.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "success", event))
}
