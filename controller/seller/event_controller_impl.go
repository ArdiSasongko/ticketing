package seller_controller

import (
	"net/http"
	"strconv"

	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/model"
	"github.com/ArdiSasongko/ticketing_app/model/web/seller"
	"github.com/ArdiSasongko/ticketing_app/service/seller"
	"github.com/labstack/echo/v4"
)

type EventControllerImpl struct {
	eventService seller_service.EventService
}

func NewEventController(eventService seller_service.EventService) *EventControllerImpl {
	return &EventControllerImpl{
		eventService: eventService,
	}
}

// List godoc
// @Summary Event (List)
// @Description Event (List)
// @Tags [Seller] Event
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
// @Failure 400 {object} helper.ResponseClientModel
// @Router /seller/events [get]
func (controller *EventControllerImpl) List(c echo.Context) error {
	filters, sort, limit, page := helper.ExtractFilterSort(c.QueryParams())

	sellerId, err := helper.GetAuthId(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseClient(http.StatusInternalServerError, err.Error(), nil))
	}

	// Call the service to fetch the list of events
	events, err := controller.eventService.GetEventList(sellerId, filters, sort, limit, page)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseClient(http.StatusInternalServerError, err.Error(), nil))
	}
	// Return JSON response with the list of events
	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success", events))
}

// View godoc
// @Summary Event (View)
// @Description Event (View)
// @Tags [Seller] Event
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param id path string true "ID"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /seller/events/{id} [get]
func (controller *EventControllerImpl) View(c echo.Context) error {
	eventId, _ := strconv.Atoi(c.Param("id"))

	event, getEventErr := controller.eventService.ViewEvent(eventId)
	if getEventErr != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, getEventErr.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Success", event))
}

// Create godoc
// @Summary Event (Create)
// @Description Event (Create)
// @Tags [Seller] Event
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param event body seller_web.CreateEventRequest true "Create Event Request"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /seller/events [post] // todo
func (controller *EventControllerImpl) Create(c echo.Context) error {
	events := new(seller_web.CreateEventRequest)

	if err := c.Bind(events); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	userID, _ := helper.GetAuthId(c)

	saveEvents, errSaveEvents := controller.eventService.SaveEvents(userID, *events)

	if errSaveEvents != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errSaveEvents.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Create Event Success", saveEvents))
}

// Update godoc
// @Summary  Event (Update)
// @Description  Event (Update)
// @Tags [Seller] Event
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param id path int true "Event ID"
// @Param event body seller_web.UpdateEventRequest true "Update Event Request"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /seller/events/{id} [put]
func (controller *EventControllerImpl) Update(c echo.Context) error {
	// Mengambil ID acara dari parameter rute
	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, "Invalid event ID", nil))
	}

	// Mengikat data dari permintaan ke struktur UpdateEventRequest
	updateRequest := new(seller_web.UpdateEventRequest)
	if err := c.Bind(updateRequest); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	// Memanggil layanan untuk melakukan pembaruan acara
	updatedEvent, err := controller.eventService.UpdateEvent(*updateRequest, eventID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	// Mengembalikan respons dengan data acara yang diperbarui
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Update Event Success", updatedEvent))
}

// UpdateStatus godoc
// @Summary  Event (Update Status)
// @Description  Event (Update Status)
// @Tags [Seller] Event
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param id path int true "Event ID"
// @Param event body seller_web.UpdateEventStatusRequest true "Update Event Request"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /seller/events/{id}/update-status [patch]
func (controller *EventControllerImpl) UpdateStatus(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, "Invalid event ID", nil))
	}

	request := new(seller_web.UpdateEventStatusRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}
	if err := c.Validate(request); err != nil {
		return err
	}

	event, err := controller.eventService.UpdateEventStatus(*request, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Update Event Status Success", event))
}

// CheckInTicket godoc
// @Summary Event (Check In Ticket)
// @Description (Check In Ticket)
// @Tags [Seller] Event
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param event_id path int true "ID"
// @Param ticket_id path int true "Ticket ID"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /seller/events/{event_id}/tickets/{ticket_id}/check-in [put] // todo
func (controller *EventControllerImpl) CheckInTicket(c echo.Context) error {
	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, "Invalid event ID", nil))
	}

	ticketID, err := strconv.Atoi(c.Param("ticket_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, "Invalid ticket ID", nil))
	}

	err = controller.eventService.CheckInTicket(eventID, ticketID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Ticket successfully checked in", nil))
}

// Delete godoc
// @Summary Event (Delete)
// @Description Event (Delete)
// @Tags [Seller] Event
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param id path int true "ID"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /seller/events/{id} [delete]
func (controller *EventControllerImpl) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if errDeleteEvent := controller.eventService.DeleteEvent(id); errDeleteEvent != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, errDeleteEvent.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Delete Event Success", nil))
}
