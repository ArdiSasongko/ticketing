package controller

type EventController interface {
	GetEventList(c echo.Context) error
	DeleteEvent(c echo.Context) error
}
