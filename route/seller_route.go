package route

import (
	"github.com/ArdiSasongko/ticketing_app/app"
	"github.com/ArdiSasongko/ticketing_app/controller/seller"
	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/middleware"
	"github.com/ArdiSasongko/ticketing_app/query_builder/seller"
	"github.com/ArdiSasongko/ticketing_app/repository/seller"
	"github.com/ArdiSasongko/ticketing_app/service/seller"
	"github.com/labstack/echo/v4"
)

// RegisterSellerRoutes mendaftarkan rute-rute untuk seller
// @Summary Register Seller Routes
// @Description Register all the seller related routes
// @Tags seller

func RegisterSellerRoutes(prefix string, e *echo.Echo) {
	db := app.DBConnection()
	token := helper.NewTokenUseCase()
	sellerAuthRepo := seller_repository.NewAuthRepository(db)
	sellerAuthService := seller_service.NewAuthService(sellerAuthRepo, token)
	sellerAuthController := seller_controller.NewAuthController(sellerAuthService)
	sellerEventQB := seller_query_builder.NewEventQueryBuilder(db)
	sellerEventRepo := seller_repository.NewEventRepository(sellerEventQB)
	sellerEventService := seller_service.NewEventService(sellerEventRepo)
	sellerEventController := seller_controller.NewEventController(sellerEventService)

	g := e.Group(prefix)

	authRoute := g.Group("/auth")
	authRoute.POST("/register", sellerAuthController.Register)
	authRoute.POST("/login", sellerAuthController.Login)

	g.Use(middleware.JWTProtection(), middleware.IsSeller)

	meRoute := g.Group("/me")
	meRoute.GET("", sellerAuthController.View)
	meRoute.PUT("/update", sellerAuthController.Update)

	eventRoute := g.Group("/events")
	eventRoute.GET("", sellerEventController.List)
	eventRoute.GET("/:id", sellerEventController.View, middleware.AccessEvent(*sellerEventRepo))
	eventRoute.POST("", sellerEventController.Create)
	eventRoute.PUT("/:id", sellerEventController.Update, middleware.AccessEvent(*sellerEventRepo))
	eventRoute.PUT("/:id/update-status", sellerEventController.UpdateStatus, middleware.AccessEvent(*sellerEventRepo))
	eventRoute.PUT("/:id/tickets/:ticket_id/check-in", sellerEventController.CheckInTicket, middleware.AccessEvent(*sellerEventRepo))
	eventRoute.DELETE("/:id", sellerEventController.Delete, middleware.AccessEvent(*sellerEventRepo))
}
