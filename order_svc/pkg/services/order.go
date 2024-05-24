package services

import (
	"context"
	"net/http"
	"order_svc/pkg/client"
	"order_svc/pkg/database"
	"order_svc/pkg/models"
	"order_svc/pkg/proto"
)

type Service struct {
	H          database.DBConnection
	ProductSvc client.ProductServiceClient
	proto.UnimplementedOrderServiceServer
}

func (s *Service) CreateOrder(ctx context.Context, req *proto.CreateOrderRequest) (*proto.CreateOrderResponse, error) {
	product, err := s.ProductSvc.FindOne(req.ProductId)

	if err != nil {
		return &proto.CreateOrderResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	} else if product.Status >= http.StatusNotFound {
		return &proto.CreateOrderResponse{
			Status: product.Status,
			Error:  product.Error,
		}, nil
	} else if product.Data.Stock < req.Quantity {
		return &proto.CreateOrderResponse{
			Status: http.StatusConflict,
			Error:  "Not enough stock",
		}, nil
	}

	order := models.Order{
		ProductId: product.Data.Id,
		Price:     int64(product.Data.Price),
		UserId:    req.UserId,
	}

	s.H.Conn.Create(&order)

	res, err := s.ProductSvc.DecreaseStock(req.ProductId, order.Id)
	if err != nil {
		return &proto.CreateOrderResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	} else if res.Status == http.StatusConflict {
		s.H.Conn.Delete(&order)

		return &proto.CreateOrderResponse{
			Status: http.StatusConflict,
			Error:  res.Error,
		}, nil
	}

	return &proto.CreateOrderResponse{
		Status: http.StatusCreated,
		Id:     order.Id,
	}, nil
}
