package main

import (
	"customer/app"
	"customer/config"
	"customer/handler"
	pb "customer/proto"
	"customer/repository"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "customer"
	version = "latest"
)

func main() {
	config.InitConfig()
	db := app.InitGorm()
	customerRepository := repository.NewCustomerRepository()
	customerHandler := handler.NewCustomerHandler(db, customerRepository)

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
	if err := pb.RegisterCustomerServiceHandler(srv.Server(), &customerHandler); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
