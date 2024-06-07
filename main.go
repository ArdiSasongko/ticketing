package main

import (
	"github.com/ArdiSasongko/ticketing_app/route"
	"github.com/labstack/echo/v4"
)

func main() {
	r := echo.New()

	route.RegisterBuyerRoutes("/buyer", r)

	route.RegisterSellerRoutes("/seller", r)

	r.Logger.Fatal(r.Start(":8001"))
}
