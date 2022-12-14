package main

import (
	"customer/app"
	"customer/config"
	"customer/handler"
	pb "customer/proto"
	"customer/repository"
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	version = "latest"
)

func main() {
	// log
	file, _ := os.OpenFile(fmt.Sprintf("logs/app_%s.log", time.Now().Format("2006_01_02")), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logrus.SetOutput(file)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// init config
	config.InitConfig()

	//
	db := app.InitGorm()
	customerRepository := repository.NewCustomerRepository()
	customerHandler := handler.NewCustomerHandler(db, customerRepository)

	// Create service
	srv := micro.NewService(
		micro.Server(grpcs.NewServer()),
		micro.Client(grpcc.NewClient()),
	)
	srv.Init(
		micro.Name(config.Config.Server.ServiceName),
		micro.Version(version),
		micro.Address(config.Config.Server.HostPort),
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
