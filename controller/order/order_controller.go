package order_controller

import "github.com/labstack/echo/v4"

type OrderController interface {
	Order(c echo.Context) error
	GenerateTicket(c echo.Context) error
	GenerateHistory(c echo.Context) error
	GenerateHistoryItem(c echo.Context) error
	GetByID(c echo.Context) error
}
