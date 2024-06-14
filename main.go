package main

import (
	"fmt"
	"github.com/ArdiSasongko/ticketing_app/route"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"os"

	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/ArdiSasongko/ticketing_app/docs" // Import the generated documentation files
)

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

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("error loading .env file!")
	}

	r := echo.New()

	// Rute Swagger
	r.GET("/swagger/*", echoSwagger.WrapHandler)

	// Daftarkan rute
	route.RegisterBuyerRoutes("/buyer", r)
	route.RegisterSellerRoutes("/seller", r)

	r.Debug = true
	r.Logger.Fatal(r.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
