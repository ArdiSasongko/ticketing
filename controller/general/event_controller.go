package general_controller

import "github.com/labstack/echo/v4"

type EventController interface {
	List(c echo.Context) error
	View(c echo.Context) error
}
