package handler

import (
	"e-commerce/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	CartService *service.CartService
}

func NewCartHandler(s *service.CartService) *CartHandler {
	return &CartHandler{CartService: s}
}

func (h *CartHandler) AddToCart(c *gin.Context) {

	userID := c.GetUint("userID")

	var req struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.CartService.AddToCart(userID, req.ProductID, req.Quantity)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "added to cart"})
}

func (h *CartHandler) GetCart(c *gin.Context) {

	userID := c.GetUint("userID")

	carts, err := h.CartService.GetCart(userID)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, carts)
}

func (h *CartHandler) RemoveCart(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	err := h.CartService.Remove(uint(id))

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "removed"})
}


func (h  *CartHandler)UpdateCart(c *gin.Context){
	cartIDParam :=c.Param("id")

	cartID,err:=strconv.Atoi(cartIDParam)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"invalid cart id"})
		return
	}
	var  req struct {
		Quantity int ` json:"quantity"`
	}
	if err :=c.ShouldBindJSON(&req);err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	err = h.CartService.UpdateQuantity(uint(cartID),req.Quantity)

	if err !=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"cart quantity updated",
	})
}