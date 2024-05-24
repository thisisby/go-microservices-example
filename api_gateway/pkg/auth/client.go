package auth

import (
	"api_gateway/pkg/auth/proto"
	"api_gateway/pkg/config"
	"google.golang.org/grpc"
	"log"
)

type ServiceClient struct {
	Client proto.AuthServiceClient
}

func InitServiceClient(c *config.Config) proto.AuthServiceClient {
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	return proto.NewAuthServiceClient(cc)
}
