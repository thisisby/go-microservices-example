package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"order_svc/pkg/client"
	"order_svc/pkg/config"
	"order_svc/pkg/database"
	"order_svc/pkg/proto"
	"order_svc/pkg/services"
)

func main() {
	c := config.LoadConfig()

	db := database.InitializeDBConnection(c.DSN)

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	productSvc := client.InitProductServiceClient(c.ProductSvcUrl)

	log.Printf("Listening on %s", c.Port)

	s := services.Service{
		H:          db,
		ProductSvc: productSvc,
	}

	grpcServer := grpc.NewServer()

	proto.RegisterOrderServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
