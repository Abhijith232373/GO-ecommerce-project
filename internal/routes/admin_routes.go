package routes

import (
	"e-commerce/internal/handler"
	"e-commerce/internal/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.Engine,adminHandler *handler.AdminHandler){
	admin:=r.Group("/admin")

	// Public routes
	admin.GET("/login",adminHandler.ShowLoginPage)
	admin.POST("/login",adminHandler.Login)

	//Protected routes
	admin.Use(middleware.AdminAuthMiddleware())
	{

		admin.GET("/dashboard",adminHandler.ShowDashboard)
		admin.GET("/products",adminHandler.ShowProductsPage)
		admin.GET("/logout",adminHandler.Logout)

//////////////// User management ///////////////////
		admin.POST("/users/add",adminHandler.AddUser)
		admin.POST("/users/suspend/:id",adminHandler.SuspendUser)
		admin.GET("/users",adminHandler.ShowUserPage)
		admin.POST("/users/activate/:id",adminHandler.ActiveUser)

/////////////// Product management /////////////////////
products := admin.Group("/products")
{

	products.GET("/products",adminHandler.ShowProductsPage)
	products.POST("add",adminHandler.AddProduuct)
	products.POST("activate/:id",adminHandler.ActivateProduct)
	products.POST("suspend/:id",adminHandler.SuspendProduct)
	products.POST("delete/:id",adminHandler.DeleteProduct)
}
}
	
}