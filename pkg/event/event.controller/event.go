package eventcontroller

import "github.com/labstack/echo/v4"

type EventController interface {
	SaveEvents(c echo.Context) error
	UpdateEvent(c echo.Context) error
}