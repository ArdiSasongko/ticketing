package seller

import "github.com/labstack/echo/v4"

type SellerControler interface {
	UpdateSeller(c echo.Context) error
}
