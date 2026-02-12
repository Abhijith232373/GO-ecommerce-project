package handler

import (
	"e-commerce/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct{
	ProductService *service.ProductService
}

func NewProductHandler(S *service.ProductService)*ProductHandler{
	return &ProductHandler{ProductService: S}
}

func (h *ProductHandler)GetProducts(c *gin.Context){
	search:=c.Query("search")
	category:=c.Query("category")

	page,_:=strconv.Atoi(c.DefaultQuery("page","1"))
	limit,_:=strconv.Atoi(c.DefaultQuery("limit","10"))

	products,total,err:=h.ProductService.GetProducts(
		search,
		category,
		page,
		limit,
	)
	if err !=nil {
		c.JSON(500,gin.H{"error":err.Error()})
		return
	}
	c.JSON(200,gin.H{
		"total":total,
		"page":page,
		"limit":limit,
		"products":products,
	})
}



