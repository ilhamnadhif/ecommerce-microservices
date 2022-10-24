package main

import (
	"cart/app"
	"cart/config"
	"cart/handler"
	pb "cart/proto"
	"cart/repository"
	"fmt"
	"os"
	"time"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
	"github.com/sirupsen/logrus"
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
	cartRepository := repository.NewCartRepository()
	cartHandler := handler.NewCartHandler(db, cartRepository)

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
	if err := pb.RegisterCartServiceHandler(srv.Server(), &cartHandler); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
