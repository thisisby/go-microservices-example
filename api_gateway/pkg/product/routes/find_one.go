package routes

import (
	"api_gateway/pkg/product/proto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FindOne(ctx *gin.Context, c proto.ProductServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.FindOne(ctx, &proto.FindOneRequest{
		Id: id,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
