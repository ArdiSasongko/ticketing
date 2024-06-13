package buyer_controller

import "github.com/labstack/echo/v4"

type OrderController interface {
	ListOrder(c echo.Context) error
	ViewOrder(c echo.Context) error
	CreateOrder(c echo.Context) error
	PayOrder(c echo.Context) error
	DeleteOrder(c echo.Context) error
}
