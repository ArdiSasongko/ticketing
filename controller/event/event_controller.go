package controller

import "github.com/labstack/echo"

type EventController interface {
	GetEventList(c echo.Context) error
	DeleteEvent(c echo.Context) error
}
