package main

import (
	"context"
	"google.golang.org/grpc/status"
	"log"
	"time"

	pb "product-client/proto"

	"github.com/go-micro/plugins/v4/client/grpc"
	"go-micro.dev/v4"
)

var (
	service = "product"
	version = "latest"
)

func main() {
	// Create service

	srv := micro.NewService(
		micro.Client(grpc.NewClient()),
	)

	srv.Init()

	// Create client
	c := pb.NewProductService(service, srv.Client())

	// Call service
	_, err := c.FindOneByID(context.Background(), &pb.ProductID{ID: 12})
	e, ok := status.FromError(err)
	if ok {
		log.Println(e.Message())
		log.Fatalln(e.Code())
	} else {
		log.Fatalln(err)
	}

	//logger.Info(rsp)

	time.Sleep(1 * time.Second)
}
