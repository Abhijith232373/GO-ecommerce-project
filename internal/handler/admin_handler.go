package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct{}

func NewAdminHandler() *AdminHandler {
	return &AdminHandler{}
}

func (h *AdminHandler) LoginPage(c *gin.Context){
	c.HTML(http.StatusOK,"admin/login.html",nil)
}

func (h *AdminHandler)DashboardPage(c *gin.Context){
	c.HTML(http.StatusOK,"admin/dashboard.html",gin.H{
		"title":"Admin Dashboard",
	})
}
func (h *AdminHandler)ProductsPage(c *gin.Context){
	c.HTML(http.StatusOK,"admin/products.html",gin.H{
		"tilte":"Products",
	})
}