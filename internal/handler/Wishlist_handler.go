package handler

import (
	"e-commerce/internal/service"
	"strconv"
	"github.com/gin-gonic/gin"
)

type WishlistHandler struct {
	Service *service.WishlistService
}

func NewWishlistHandler(s *service.WishlistService)*WishlistHandler{
	return &WishlistHandler{Service: s}
}

func (h *WishlistHandler)Add(c *gin.Context){
	userID:=c.GetUint("userID")

	var req struct{
		ProductID uint `json:"product_id"`
	}
	if err:=c.ShouldBindJSON(&req);err !=nil{
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}
	err:=h.Service.Add(userID,req.ProductID)

	if err !=nil{
		c.JSON(500,gin.H{"error":err.Error()})
		return
	}
	c.JSON(200,gin.H{"message":"added to wishlist"})
}

func (h *WishlistHandler)Get(c *gin.Context){
	userID :=c.GetUint("userID")

	data,err:=h.Service.Get(userID)

	if err!=nil{
		c.JSON(500,gin.H{"error":err.Error()})
		return
	}
	c.JSON(200,data)
}

func (h *WishlistHandler)Remove(c *gin.Context){
	id,_:=strconv.Atoi(c.Param("id"))

	err:=h.Service.Remove(uint(id))

	if err !=nil{
		c.JSON(500,gin.H{"error":err.Error()})
		return
	}
	c.JSON(200,gin.H{"message":"removed from wishlist..."})
}