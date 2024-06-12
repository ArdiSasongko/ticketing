package route

import (
	"github.com/ArdiSasongko/ticketing_app/app"
	seller_controller "github.com/ArdiSasongko/ticketing_app/controller/seller"
	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/middleware"
	seller_query_builder "github.com/ArdiSasongko/ticketing_app/query_builder/seller"
	seller_repository "github.com/ArdiSasongko/ticketing_app/repository/seller"
	seller_service "github.com/ArdiSasongko/ticketing_app/service/seller"
	"github.com/labstack/echo/v4"
)

func RegisterSellerRoutes(prefix string, e *echo.Echo) {
	db := app.DBConnection()
	token := helper.NewTokenUseCase()
	sellerAuthRepo := seller_repository.NewSellerRepository(db)
	sellerAuthService := seller_service.NewSellerService(sellerAuthRepo, token)
	sellerAuthController := seller_controller.NewSellerController(sellerAuthService)
	sellerEventQB := seller_query_builder.NewEventQueryBuilder(db)
	sellerEventRepo := seller_repository.NewEventRepository(sellerEventQB)
	sellerEventService := seller_service.NewEventService(sellerEventRepo)
	sellerEventController := seller_controller.NewEventController(sellerEventService)

	g := e.Group(prefix)

	authRoute := g.Group("/auth")
	authRoute.POST("/register", sellerAuthController.SaveSeller)
	authRoute.POST("/login", sellerAuthController.LoginSeller)

	meRoute := g.Group("/me", middleware.JWTProtection())
	meRoute.POST("", sellerAuthController.GetSeller)
	meRoute.POST("/update", sellerAuthController.UpdateSeller)

	eventRoute := g.Group("/events", middleware.JWTProtection())
	eventRoute.GET("", sellerEventController.GetEventList)
	eventRoute.POST("", sellerEventController.SaveEvents)
	eventRoute.POST(":id", sellerEventController.UpdateEvent, middleware.AccessUserID(*sellerEventRepo))
}
