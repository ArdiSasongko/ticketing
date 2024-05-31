package seller

import "github.com/labstack/echo"

type UserController interface {
	GetEventList(c echo.Context) error
}
