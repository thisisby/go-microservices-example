package routes

import (
	"api_gateway/pkg/product/proto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateProductRequestBody struct {
	Name  string  `json:"name"`
	Stock int64   `json:"stock"`
	Price float32 `json:"price"`
}

func CreateProduct(ctx *gin.Context, c proto.ProductServiceClient) {
	var req CreateProductRequestBody

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateProduct(ctx, &proto.CreateProductRequest{
		Name:  req.Name,
		Stock: req.Stock,
		Price: req.Price,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
