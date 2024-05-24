package product

import (
	"api_gateway/pkg/config"
	"api_gateway/pkg/product/proto"
	"google.golang.org/grpc"
	"log"
)

type ServiceClient struct {
	Client proto.ProductServiceClient
}

func InitServiceClient(c *config.Config) proto.ProductServiceClient {
	cc, err := grpc.Dial(c.ProductSvcUrl, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect to %s: %v", c.ProductSvcUrl, err)
	}

	return proto.NewProductServiceClient(cc)
}
