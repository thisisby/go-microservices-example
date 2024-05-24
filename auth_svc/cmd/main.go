package main

import (
	"auth_svc/pkg/config"
	"auth_svc/pkg/database"
	"auth_svc/pkg/proto"
	"auth_svc/pkg/services"
	"auth_svc/pkg/utils"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	c := config.LoadConfig()

	db := database.InitializeDBConnection(c.DSN)

	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Listening on %s", c.Port)

	s := services.Server{
		H:   db,
		Jwt: jwt,
	}

	grpcServer := grpc.NewServer()

	proto.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}

}
