package routes

import (
	"e-commerce/internal/handler"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.Engine,adminHandler *handler.AdminHandler){
	admin:=r.Group("/admin")
	admin.GET("/")
}