package seller_controller

import "github.com/labstack/echo/v4"

type UserController interface {
	GetEventList(c echo.Context) error
	ViewEvent(c echo.Context) error
	SaveEvents(c echo.Context) error
	UpdateEvent(c echo.Context) error
	CheckInTicket(c echo.Context) error
	DeleteEvent(c echo.Context) error
}
