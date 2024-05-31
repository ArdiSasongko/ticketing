package seller

import "github.com/labstack/echo/v4"

type SellerControler interface {
	GetSeller(c echo.Context) error
	UpdateSeller(c echo.Context) error
}
