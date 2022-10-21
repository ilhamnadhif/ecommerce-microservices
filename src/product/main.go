package main

import (
	"fmt"
	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
	"github.com/sirupsen/logrus"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"os"
	"product/app"
	"product/config"
	"product/handler"
	pb "product/proto"
	"product/repository"
	"time"
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
	productRepository := repository.NewProductRepository()
	productHandler := handler.NewProductHandler(db, productRepository)

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
	if err := pb.RegisterProductServiceHandler(srv.Server(), &productHandler); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
