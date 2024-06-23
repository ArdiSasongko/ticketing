package seller_controller

import "github.com/labstack/echo/v4"

type UserController interface {
	List(c echo.Context) error
	View(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	UpdateStatus(c echo.Context) error
	CheckInTicket(c echo.Context) error
	Delete(c echo.Context) error
}
