package main

import (
	"fmt"
	"github.com/ArdiSasongko/ticketing_app/route"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("error loading .env file!")
	}

	r := echo.New()

	route.RegisterBuyerRoutes("/buyer", r)
	route.RegisterSellerRoutes("/seller", r)

	r.Debug = true
	r.Logger.Fatal(r.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
