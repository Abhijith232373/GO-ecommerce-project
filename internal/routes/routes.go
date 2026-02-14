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
	cartHandler *handler.CartHandler,
	wishlistHandler *handler.WishlistHandler,
	orderHandler *handler.OrderHandler,
	){


/////////////////////////  Public auth route ///////////////////////////


	auth:=r.Group("/auth")
	{
		auth.POST("/",userHandler.Home)
		auth.POST("/register",userHandler.Register)
		auth.POST("/login",userHandler.Login)
	}  

/////////////////////// PRODUCT --////////////////////////////////////

		r.GET("/products",productHandler.GetProducts)
		r.GET("/products/:id",productHandler.GetProducts)

		
		
//////////////////////// protected route ////////////////////////
		
		api:=r.Group("/api")
		api.Use(middleware.JWTAuth())
		{
			api.GET("/profile",userHandler.Profile)
			api.POST("/logout",userHandler.Logout)

			
//////////////////////// cart ////////////////////////////////////////

			api.POST("/cart", cartHandler.AddToCart)
			api.GET("/cart", cartHandler.GetCart)
			api.DELETE("/cart/:id", cartHandler.RemoveCart)
			api.PUT("/cart/:id",cartHandler.UpdateCart)

////////////////////// WishlistHandler ///////////////////////////////

			api.POST("/wishlist",wishlistHandler.Add)
			api.GET("/wishlist",wishlistHandler.Get)
			api.DELETE("/wishlist/:id",wishlistHandler.Remove)

///////////////////////// Order  ////////////////////////////////////////

			api.POST("/orders",orderHandler.CreateOrder)
			api.GET("/orders",orderHandler.GetOrders)
			api.GET("/orders/:id",orderHandler.GetOrderByID)
		}
}