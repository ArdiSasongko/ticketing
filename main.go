package main

import (
	"github.com/ArdiSasongko/ticketing_app/app"
	sellercontroller "github.com/ArdiSasongko/ticketing_app/controller/seller"
	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/repository/seller"
	sellerservice "github.com/ArdiSasongko/ticketing_app/service/seller"
	"github.com/labstack/echo/v4"
)

func main() {
	r := echo.New()
	db := app.DBConnection()
	token := helper.NewTokenUseCase()
	Sellerrepo := seller.NewSellerRepository(db)
	SellerService := sellerservice.NewSellerService(Sellerrepo, token)
	sellercontroller := sellercontroller.NewSellerController(SellerService)

	r.POST("/register/seller", sellercontroller.SaveSeller)
	r.POST("/login/seller", sellercontroller.LoginSeller)

	r.Logger.Fatal(r.Start(":8001"))
}
