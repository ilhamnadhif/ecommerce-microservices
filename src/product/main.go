package main

import (
	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"product/app"
	"product/config"
	"product/handler"
	pb "product/proto"
	"product/repository"
)

var (
	service = "product"
	version = "latest"
)

func main() {
	config.InitConfig()
	db := app.InitGorm()
	productRepository := repository.NewProductRepository()
	productHandler := handler.NewProductHandler(db, productRepository)

	// Create service
	srv := micro.NewService(
		micro.Server(grpcs.NewServer()),
		micro.Client(grpcc.NewClient()),
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	// Register handler
	if err := pb.RegisterProductServiceHandler(srv.Server(), &productHandler); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
