package handler

import (
	"e-commerce/internal/service"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	OrderService *service.OrderService
}
func NewOrderHandler(s *service.OrderService)*OrderHandler{
	return &OrderHandler{OrderService: s}
}

func (h *OrderHandler)CreatedOrder(c *gin.Context){
	userID :=c.GetUint("userID")
	
	err:=h.OrderService.CreateOrder(userID)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"order created successfully",
	})
}

func (h *OrderHandler)GetOrders(c *gin.Context){
	userID :=c.GetUint("userID")
	orders,err:=h.OrderService.GetUserOrders(userID)

	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,orders)
}


func (h *OrderHandler)GetOrderByID(c *gin.Context){
	idParam:=c.Param("id")

	orderID,err:=strconv.Atoi(idParam)

	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"invalid order id",
		})
		return
	}
	order,err:=h.OrderService.GetOrderByID(uint(orderID))

	if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{
			"error":"order not found",
		})
		return
	}
	c.JSON(http.StatusOK,order)
	}

func (h *OrderHandler) CreateOrder(c *gin.Context) {

	userID := c.GetUint("userID")

	var req struct {
		FullName string `json:"full_name"`
		Phone    string `json:"phone"`
		Address  string `json:"address"`
		City     string `json:"city"`
		State    string `json:"state"`
		Pincode  string `json:"pincode"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
		addr := service.AddressInput{
		FullName: req.FullName,
		Phone:    req.Phone,
		Address:  req.Address,
		City:     req.City,
		State:    req.State,
		Pincode:  req.Pincode,
	}
	err := h.OrderService.CreateOrderWithAddress(userID, addr)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "order created successfully"})
}