package buyer

import "github.com/labstack/echo/v4"

type BuyerControllerInterface interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
}
