package buyer_controller

import "github.com/labstack/echo/v4"

type EventController interface {
	GetEventList(c echo.Context) error
	ViewEvent(c echo.Context) error
}
