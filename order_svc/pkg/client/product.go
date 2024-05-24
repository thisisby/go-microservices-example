package client

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"order_svc/pkg/proto"
)

type ProductServiceClient struct {
	Client proto.ProductServiceClient
}

func InitProductServiceClient(url string) ProductServiceClient {
	cc, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Printf("Error dialing to product services: %v", err)
	}

	c := ProductServiceClient{
		Client: proto.NewProductServiceClient(cc),
	}

	return c
}

func (c *ProductServiceClient) FindOne(id int64) (*proto.FindOneResponse, error) {
	req := &proto.FindOneRequest{
		Id: id,
	}

	return c.Client.FindOne(context.Background(), req)
}

func (c *ProductServiceClient) DecreaseStock(id int64, orderId int64) (*proto.DecreaseStockResponse, error) {
	req := &proto.DecreaseStockRequest{
		Id:      id,
		OrderId: orderId,
	}

	return c.Client.DecreaseStock(context.Background(), req)
}
