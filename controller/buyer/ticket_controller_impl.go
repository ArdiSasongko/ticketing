package buyer_controller

import (
	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/model"
	"github.com/ArdiSasongko/ticketing_app/service/buyer"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type TicketControllerImpl struct {
	ticketService buyer_service.TicketService
}

func NewTicketController(ticketService buyer_service.TicketService) *TicketControllerImpl {
	return &TicketControllerImpl{
		ticketService: ticketService,
	}
}

// List godoc
// @Summary Ticket (List)
// @Description Ticket (List)
// @Tags [Buyer] Ticket
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param filter[status] query string false "Filter"
// @Param sort query string false "Sort"
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 500 {object} helper.ResponseClientModel
// @Router /buyer/tickets [get]
func (controller *TicketControllerImpl) List(c echo.Context) error {
	filters, sort, limit, page := helper.ExtractFilterSort(c.QueryParams())
	histories, err := controller.ticketService.GetTicketList(filters, sort, limit, page)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Success", histories))
}

// View godoc
// @Summary Ticket (View)
// @Description Ticket (View)
// @Tags [Buyer] Ticket
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param id path int true "id"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 500 {object} helper.ResponseClientModel
// @Router /buyer/tickets/{id} [get]
func (controller *TicketControllerImpl) View(c echo.Context) error {
	ticketId, _ := strconv.Atoi(c.Param("id"))

	history, err := controller.ticketService.ViewTicket(ticketId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Success", history))
}
