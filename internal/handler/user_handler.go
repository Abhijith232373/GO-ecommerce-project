package handler

import (
	"e-commerce/internal/models"
	"e-commerce/internal/service"
	"net/http"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{UserService: s}
}

func (h *UserHandler)Home(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{"message":"welcome home page !"})
}

///////////////////////// signup//////////////////

func (h *UserHandler) Register(c *gin.Context) {
	var user models.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.UserService.Register(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "registered"})
}

/////////////////////  login //////////////////////////////

func (h *UserHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	access, refresh, err := h.UserService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"access_token":  access,
		"refresh_token": refresh,
	})
}

////////////////////////// Logout ///////////////////////////

func (h *UserHandler)Logout (c *gin.Context){
	userIDValue,exists :=c.Get("userID")

	if !exists{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"Unauthorized"})
		return
	}
	userID,ok:=userIDValue.(uint)
	if !ok{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"Invalid user"})
		return
	}
	if err:=h.UserService.Logout(userID);err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"Logged out successfully",
	})
}
