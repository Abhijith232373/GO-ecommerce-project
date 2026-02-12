package routes

import (
	"e-commerce/internal/handler"
	"e-commerce/internal/middleware"


	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	r *gin.Engine,
	userHandler *handler.UserHandler,
	productHandler *handler.ProductHandler,
	){


// ///////////////////////  Public auth route ///////////////////////////


	auth:=r.Group("/auth")
	{
		auth.POST("/",userHandler.Home)
		auth.POST("/register",userHandler.Register)
		auth.POST("/login",userHandler.Login)
	}  

/////////////////////// PRODUCT --////////////////////////////////////

		r.GET("/products",productHandler.GetProducts)
		r.GET("/products/:id",productHandler.GetProducts)

	// ////////////////////// protected route ////////////////////////

	api:=r.Group("/api")
	api.Use(middleware.JWTAuth())
{
	api.GET("/profile",userHandler.Profile)
	api.POST("/logout",userHandler.Logout)
}
}