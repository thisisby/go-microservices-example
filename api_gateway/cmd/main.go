package main

import (
	"api_gateway/pkg/auth"
	"api_gateway/pkg/config"
	"api_gateway/pkg/order"
	"api_gateway/pkg/product"
	"github.com/gin-gonic/gin"
)

func main() {
	c := config.LoadConfig()

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)

}
