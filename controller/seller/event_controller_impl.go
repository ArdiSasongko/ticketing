package seller_controller

import (
	"net/http"
	"strconv"

	"github.com/ArdiSasongko/ticketing_app/model"
	"github.com/ArdiSasongko/ticketing_app/model/web/seller"
	"github.com/golang-jwt/jwt/v5"

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

// GetEventList godoc
// @Summary Get list of events for a seller
// @Description Get list of events based on seller ID with optional filters, sorting, pagination
// @Tags seller
// @Accept json
// @Produce json
// @Param seller_id query int true "Seller ID"
// @Param filters query string false "Filters for events (e.g., name=EventName)"
// @Param sort query string false "Sort order for events (e.g., +date or -name)"
// @Param limit query int false "Limit number of events per page"
// @Param page query int false "Page number"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /seller/events [get]
func (controller *EventControllerImpl) GetEventList(c echo.Context) error {
	// Extract filters, sort, limit, and page from query parameters
	filters, sort, limit, page := helper.ExtractFilterSort(c.QueryParams())

	// TODO: Implement seller authentication to get the seller ID dynamically
	sellerId := 1 // For demonstration, replace this with your seller authentication logic

	// Call the service to fetch the list of events
	events, err := controller.eventService.GetEventList(sellerId, filters, sort, limit, page)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseClient(http.StatusInternalServerError, err.Error(), nil))
	}
	// Return JSON response with the list of events
	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "success", events))
}

// SaveEvents godoc
// @Summary Create a new event
// @Description Create a new event with the input payload
// @Tags seller
// @Accept json
// @Produce json
// @Param event body seller_web.CreateEventsRequest true "Create Event Request"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /seller/events [post]
func (controller *EventControllerImpl) ViewEvent(c echo.Context) error {
	eventId, _ := strconv.Atoi(c.Param("id"))

	event, getEventErr := controller.eventService.ViewEvent(eventId)
	if getEventErr != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, getEventErr.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "view event success", event))
}

func (controller *EventControllerImpl) SaveEvents(c echo.Context) error {
	events := new(seller_web.CreateEventsRequest)

	if err := c.Bind(events); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*helper.JwtCustomClaims)
	userID, _ := strconv.Atoi(claims.ID)

	saveEvents, errSaveEvents := controller.eventService.SaveEvents(userID, *events)

	if errSaveEvents != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errSaveEvents.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Successfully created event", saveEvents))
}

// UpdateEvent godoc
// @Summary Update an event
// @Description Update an event by its ID
// @Tags seller
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Param event body seller_web.UserUpdateServiceRequest true "Update Event Request"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /seller/events/{id} [put]
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

// UpdateEventStatus godoc
// @Summary Update event status
// @Description Update event status
// @Tags seller
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Param event body seller_web.UpdateEventStatusRequest true "Update Event Request"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /seller/events/{id}/status [patch]
func (controller *EventControllerImpl) UpdateEventStatus(c echo.Context) error {
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

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Successfully updated event", event))
}

// CheckInTicket godoc
// @Summary Check in ticket
// @Description Check in ticket
// @Tags seller
// @Accept json
// @Produce json
// @Param event_id path int
// @Param ticket_id path int
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /seller/events/:event_id/tickets/:ticket_id/check-in [put]
func (controller *EventControllerImpl) CheckInTicket(c echo.Context) error {
	eventID, err := strconv.Atoi(c.Param("event_id"))
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

// DeleteEvent godoc
// @Summary Delete event
// @Description Delete event
// @Tags seller
// @Accept json
// @Produce json
// @Param id path int
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /seller/events/:id [delete]
func (controller *EventControllerImpl) DeleteEvent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if errDeleteEvent := controller.eventService.DeleteEvent(id); errDeleteEvent != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, errDeleteEvent.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Event berhasil dihapus", nil))
}
