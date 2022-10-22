package app

import (
	"api-gateway/config"
	"api-gateway/dto"
	"api-gateway/handler"
	"api-gateway/middleware"
	pb "api-gateway/proto"
	"api-gateway/service"
	"github.com/go-micro/plugins/v4/client/grpc"
	"github.com/labstack/echo/v4"
	middleware2 "github.com/labstack/echo/v4/middleware"
	"go-micro.dev/v4"
)

func Route() *echo.Echo {

	// service
	srv := micro.NewService(
		micro.Client(grpc.NewClient()),
	)
	srv.Init()

	// Create client
	productRPCClient := pb.NewProductService(config.Config.Service[config.ProductService].ServiceName, srv.Client())
	merchantRPCClient := pb.NewMerchantService(config.Config.Service[config.MerchantService].ServiceName, srv.Client())
	customerRPCClient := pb.NewCustomerService(config.Config.Service[config.CustomerService].ServiceName, srv.Client())

	// service
	productService := service.NewProductService(productRPCClient, merchantRPCClient)
	merchantService := service.NewMerchantService(merchantRPCClient, productRPCClient)
	authService := service.NewAuthService(merchantRPCClient, customerRPCClient)

	// handler
	productHandler := handler.NewProductHandler(productService)
	merchantHandler := handler.NewMerchantHandler(merchantService)
	authHandler := handler.NewAuthHandler(authService)

	config := middleware2.JWTConfig{
		Claims:     &dto.JWTCustomClaims{},
		SigningKey: []byte(config.Config.Jwt.SigningKey),
	}
	e := echo.New()
	e.HTTPErrorHandler = middleware.CustomHTTPErrorHandler
	e.Use(middleware2.CORS())
	e.Use(middleware2.Recover())

	apiRouter := e.Group("/api")

	authRouter := apiRouter.Group("/login")
	authRouter.POST("/merchant", authHandler.LoginMerchant)
	authRouter.POST("/customer", authHandler.LoginCustomer)

	apiRouter.Use(middleware2.JWTWithConfig(config))

	productRouter := apiRouter.Group("/products")
	productRouter.GET("", productHandler.FindAll)
	productRouter.GET("/:productID", productHandler.FindOneByID)
	productRouter.POST("", productHandler.Create)
	productRouter.PUT("/:productID", productHandler.Update)
	productRouter.DELETE("/:productID", productHandler.Delete)

	merchantRouter := apiRouter.Group("/merchants")
	merchantRouter.GET("", merchantHandler.FindAll)
	merchantRouter.GET("/:merchantID", merchantHandler.FindOneByID)
	merchantRouter.POST("", merchantHandler.Create)
	merchantRouter.PUT("/:merchantID", merchantHandler.Update)
	merchantRouter.DELETE("/:merchantID", merchantHandler.Delete)

	return e
}
