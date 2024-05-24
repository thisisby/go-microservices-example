package services

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"product_svc/pkg/database"
	"product_svc/pkg/models"
	"product_svc/pkg/proto"
)

type Service struct {
	H database.DBConnection
	proto.UnimplementedProductServiceServer
}

func (s *Service) CreateProduct(ctx context.Context, req *proto.CreateProductRequest) (*proto.CreateProductResponse, error) {
	log.Printf("Creating product: %v", req)

	var product models.Product
	fmt.Println("repository")
	fmt.Println(req)
	fmt.Println("----------------")
	product.Name = req.Name
	product.Stock = req.Stock
	log.Printf("Price: %v", req.Price)
	product.Price = req.Price

	result := s.H.Conn.Create(&product)
	if result.Error != nil {
		return &proto.CreateProductResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	return &proto.CreateProductResponse{
		Status: http.StatusCreated,
		Id:     product.Id,
	}, nil
}

func (s *Service) FindOne(ctx context.Context, req *proto.FindOneRequest) (*proto.FindOneResponse, error) {
	log.Printf("Finding product by id: %v", req)

	var product models.Product
	result := s.H.Conn.First(&product, req.Id)
	if result.Error != nil {
		return &proto.FindOneResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	data := &proto.FindOneData{
		Id:    product.Id,
		Name:  product.Name,
		Stock: product.Stock,
		Price: float32(product.Price),
	}

	return &proto.FindOneResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}

func (s *Service) FindAll(ctx context.Context, req *proto.FindAllRequest) (*proto.FindAllResponse, error) {
	log.Printf("Finding all products")

	var products []models.Product
	rows, err := s.H.Conn.Model(&models.Product{}).Where("stock > 0").Rows()
	defer rows.Close()
	if err != nil {
		return &proto.FindAllResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	for rows.Next() {
		var product models.Product
		err := s.H.Conn.ScanRows(rows, &product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	var data []*proto.FindOneData
	for _, product := range products {
		data = append(data, &proto.FindOneData{
			Id:    product.Id,
			Name:  product.Name,
			Stock: product.Stock,
			Price: float32(product.Price),
		})
	}

	return &proto.FindAllResponse{
		Status:   http.StatusOK,
		Products: data,
	}, nil
}

func (s *Service) DecreaseStock(ctx context.Context, req *proto.DecreaseStockRequest) (*proto.DecreaseStockResponse, error) {
	log.Printf("Decreasing stock: %v", req)

	var product models.Product
	result := s.H.Conn.First(&product, req.Id)
	if result.Error != nil {
		return &proto.DecreaseStockResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	if product.Stock <= 0 {
		return &proto.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error:  "Not enough stock",
		}, nil
	}

	product.Stock -= 1
	result = s.H.Conn.Save(&product)
	if result.Error != nil {
		return &proto.DecreaseStockResponse{
			Status: http.StatusInternalServerError,
			Error:  result.Error.Error(),
		}, nil
	}

	stockLog := models.StockDecreaseLog{
		ProductRefer: product.Id,
	}
	result = s.H.Conn.Create(&stockLog)
	if result.Error != nil {
		return &proto.DecreaseStockResponse{
			Status: http.StatusInternalServerError,
			Error:  result.Error.Error(),
		}, nil
	}

	return &proto.DecreaseStockResponse{
		Status: http.StatusOK,
	}, nil
}
