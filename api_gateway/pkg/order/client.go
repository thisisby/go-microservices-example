package order

import (
	"api_gateway/pkg/config"
	"api_gateway/pkg/order/proto"
	"fmt"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client proto.OrderServiceClient
}

func InitServiceClient(c *config.Config) proto.OrderServiceClient {
	cc, err := grpc.Dial(c.OrderSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return proto.NewOrderServiceClient(cc)
}
