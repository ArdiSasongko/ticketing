package buyer_controller

import "github.com/labstack/echo/v4"

type TicketController interface {
	List(c echo.Context) error
	View(c echo.Context) error
}
