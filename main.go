package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/route"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cV *CustomValidator) Validate(i interface{}) error {
	return cV.validator.Struct(i)
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("error loading .env file!")
	}

	r := echo.New()
	r.Validator = &CustomValidator{validator: validator.New()}
	r.HTTPErrorHandler = helper.BindAndValidate

	route.RegisterBuyerRoutes("/buyer", r)

	route.RegisterSellerRoutes("/seller", r)

	r.Debug = true
	r.Logger.Fatal(r.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
