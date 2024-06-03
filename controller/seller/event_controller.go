package seller_controller

import "github.com/labstack/echo/v4"

type UserController interface {
	GetEventList(c echo.Context) error
}
