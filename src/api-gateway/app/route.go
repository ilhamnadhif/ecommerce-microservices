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
	cartRPCClient := pb.NewCartService(config.Config.Service[config.CartService].ServiceName, srv.Client())

	// service
	productService := service.NewProductService(productRPCClient, merchantRPCClient)
	merchantService := service.NewMerchantService(merchantRPCClient, productRPCClient)
	customerService := service.NewCustomerService(customerRPCClient, cartRPCClient, productRPCClient)
	cartService := service.NewCartService(cartRPCClient, productRPCClient)
	authService := service.NewAuthService(merchantRPCClient, customerRPCClient)

	// handler
	productHandler := handler.NewProductHandler(productService)
	merchantHandler := handler.NewMerchantHandler(merchantService)
	customerHandler := handler.NewCustomerHandler(customerService)
	cartHandler := handler.NewCartHandler(cartService)
	authHandler := handler.NewAuthHandler(authService)

	config := middleware2.JWTConfig{
		Claims:     &dto.JWTCustomClaims{},
		SigningKey: []byte(config.Config.Jwt.SigningKey),
	}
	jwtMiddleware := middleware2.JWTWithConfig(config)

	e := echo.New()
	e.HTTPErrorHandler = middleware.CustomHTTPErrorHandler
	e.Use(middleware2.CORS())
	e.Use(middleware2.Recover())

	apiRouter := e.Group("/api")

	authRouter := apiRouter.Group("/login")
	authRouter.POST("/merchant", authHandler.LoginMerchant)
	authRouter.POST("/customer", authHandler.LoginCustomer)

	productRouter := apiRouter.Group("/products")
	productRouter.GET("", productHandler.FindAll)
	productRouter.GET("/:productID", productHandler.FindOneByID)
	productRouter.POST("", productHandler.Create, jwtMiddleware)
	productRouter.PUT("/:productID", productHandler.Update, jwtMiddleware)
	productRouter.DELETE("/:productID", productHandler.Delete, jwtMiddleware)

	merchantRouter := apiRouter.Group("/merchants")
	merchantRouter.GET("", merchantHandler.FindAll)
	merchantRouter.GET("/:merchantID", merchantHandler.FindOneByID)
	merchantRouter.POST("", merchantHandler.Create)
	merchantRouter.GET("/common", merchantHandler.FindOneByCommon, jwtMiddleware)
	merchantRouter.PUT("/:merchantID", merchantHandler.Update, jwtMiddleware)
	merchantRouter.DELETE("/:merchantID", merchantHandler.Delete, jwtMiddleware)

	customerRouter := apiRouter.Group("/customers")
	customerRouter.GET("", customerHandler.FindAll)
	customerRouter.GET("/:customerID", customerHandler.FindOneByID)
	customerRouter.POST("", customerHandler.Create)
	customerRouter.GET("/common", customerHandler.FindOneByCommon, jwtMiddleware)
	customerRouter.PUT("/:customerID", customerHandler.Update, jwtMiddleware)
	customerRouter.DELETE("/:customerID", customerHandler.Delete, jwtMiddleware)

	cartRouter := apiRouter.Group("/carts")
	cartRouter.GET("", cartHandler.FindAll)
	cartRouter.GET("/:cartID", cartHandler.FindOneByID)
	cartRouter.POST("", cartHandler.Create, jwtMiddleware)
	cartRouter.PUT("/:cartID", cartHandler.Update, jwtMiddleware)
	cartRouter.DELETE("/:cartID", cartHandler.Delete, jwtMiddleware)

	return e
}
