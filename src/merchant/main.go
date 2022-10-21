package main

import (
	"merchant/app"
	"merchant/config"
	"merchant/handler"
	pb "merchant/proto"
	"merchant/repository"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "merchant"
	version = "latest"
)

func main() {
	config.InitConfig()
	db := app.InitGorm()
	merchantRepository := repository.NewMerchantRepository()
	merchantHandler := handler.NewMerchantHandler(db, merchantRepository)

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
	if err := pb.RegisterMerchantServiceHandler(srv.Server(), &merchantHandler); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
