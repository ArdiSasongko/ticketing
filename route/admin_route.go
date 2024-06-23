package route

import (
	"github.com/ArdiSasongko/ticketing_app/app"
	"github.com/ArdiSasongko/ticketing_app/controller/admin"
	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/middleware"
	"github.com/ArdiSasongko/ticketing_app/query_builder/admin"
	"github.com/ArdiSasongko/ticketing_app/repository/admin"
	"github.com/ArdiSasongko/ticketing_app/service/admin"
	"github.com/labstack/echo/v4"
)

func RegisterAdminRoute(prefix string, e *echo.Echo) {
	db := app.DBConnection()
	token := helper.NewTokenUseCase()
	adminAuthRepo := admin_repository.NewAuthRepository(db)
	adminAuthService := admin_service.NewAuthService(adminAuthRepo, token)
	adminAuthController := admin_controller.NewAuthController(adminAuthService)
	adminAdminQueryBuilder := admin_query_builder.NewAdminQueryBuilder(db)
	adminAdminRepo := admin_repository.NewAdminRepository(adminAdminQueryBuilder, db)
	adminAdminService := admin_service.NewAdminService(adminAdminRepo, token)
	adminAdminController := admin_controller.NewAdminController(adminAdminService)
	adminBuyerQueryBuilder := admin_query_builder.NewBuyerQueryBuilder(db)
	adminBuyerRepo := admin_repository.NewBuyerRepository(adminBuyerQueryBuilder, db)
	adminBuyerService := admin_service.NewBuyerService(adminBuyerRepo, token)
	adminBuyerController := admin_controller.NewBuyerController(adminBuyerService)
	adminSellerQueryBuilder := admin_query_builder.NewSellerQueryBuilder(db)
	adminSellerRepo := admin_repository.NewSellerRepository(adminSellerQueryBuilder, db)
	adminSellerService := admin_service.NewSellerService(adminSellerRepo, token)
	adminSellerController := admin_controller.NewSellerController(adminSellerService)

	g := e.Group(prefix)

	authRoute := g.Group("/auth")
	authRoute.POST("/register", adminAuthController.Register)
	authRoute.POST("/login", adminAuthController.Login)

	g.Use(middleware.JWTProtection(), middleware.IsAdmin)

	meRoute := g.Group("/me")
	meRoute.GET("", adminAuthController.View)
	meRoute.PUT("/update", adminAuthController.Update)

	adminRoute := g.Group("/admins")
	adminRoute.GET("", adminAdminController.List)
	adminRoute.GET("/:id", adminAdminController.View)

	buyerRoute := g.Group("/buyers")
	buyerRoute.GET("", adminBuyerController.List)
	buyerRoute.GET("/:id", adminBuyerController.View)

	sellerRoute := g.Group("/sellers")
	sellerRoute.GET("", adminSellerController.List)
	sellerRoute.GET("/:id", adminSellerController.View)
}
