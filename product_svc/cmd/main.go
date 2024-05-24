package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"product_svc/pkg/config"
	"product_svc/pkg/database"
	"product_svc/pkg/proto"
	"product_svc/pkg/services"
)

func main() {
	c := config.LoadConfig()

	db := database.InitializeDBConnection(c.DSN)

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Listening on %s", c.Port)

	// Create a new gRPC server
	s := services.Service{
		H: db,
	}

	// Register the services with the server
	grpcServer := grpc.NewServer()

	proto.RegisterProductServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
