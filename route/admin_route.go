package route

import (
	"github.com/ArdiSasongko/ticketing_app/app"
	admin_controller "github.com/ArdiSasongko/ticketing_app/controller/admin"
	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/middleware"
	admin_repository "github.com/ArdiSasongko/ticketing_app/repository/admin"
	admin_service "github.com/ArdiSasongko/ticketing_app/service/admin"
	"github.com/labstack/echo/v4"
)

func RegisterAdminRoute(prefix string, e *echo.Echo) {
	db := app.DBConnection()
	token := helper.NewTokenUseCase()
	adminRepo := admin_repository.NewAdminRepository(db)
	adminService := admin_service.NewAdminService(adminRepo, token)
	adminController := admin_controller.NewAdminController(adminService)

	g := e.Group(prefix)

	authRoute := g.Group("/auth")
	authRoute.POST("/register", adminController.Register)
	authRoute.POST("/login", adminController.Login)

	adminRoute := g.Group("/", middleware.JWTProtection())
	adminRoute.GET("/buyers", adminController.GetBuyers)
	adminRoute.GET("/sellers", adminController.GetSellers)
	adminRoute.GET("/buyer/:id", adminController.GetBuyerByID)
	adminRoute.GET("/seller/:id", adminController.GetSellerByID)
}
