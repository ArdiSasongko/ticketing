package main

import (
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"

	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/route"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"github.com/swaggo/echo-swagger"

	_ "github.com/ArdiSasongko/ticketing_app/docs" // Import the generated documentation files
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cV *CustomValidator) Validate(i interface{}) error {
	return cV.validator.Struct(i)
}

// @title Ticketing API
// @version 1.0
// @description This is a sample Ticketing server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8001
// @BasePath /

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Bearer

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("error loading .env file!")
	}

	r := echo.New()
	r.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
	}))

	r.Validator = helper.NewCustomValidator()
	r.HTTPErrorHandler = helper.BindAndValidate

	r.GET("/swagger/*", echoSwagger.WrapHandler)

	route.RegisterGeneralRoutes("/general", r)
	route.RegisterAdminRoute("/admin", r)
	route.RegisterBuyerRoutes("/buyer", r)
	route.RegisterSellerRoutes("/seller", r)

	r.Logger.Fatal(r.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
