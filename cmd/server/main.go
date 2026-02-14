package main

import (
	"e-commerce/internal/config"
	"e-commerce/internal/handler"
	"e-commerce/internal/repository"
	"e-commerce/internal/routes"
	"e-commerce/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
config.LoadEnv()

db:=config.ConnectDB()

//////////////////   USER    ?? ///////////////////////////////////

userRepo:=repository.NewUserRepository(db)
userService:=service.NewUserService(userRepo)
userHandler:=handler.NewUserHandler(userService)

//////////////////  Product ///////////////////////////////////////
productRepo:=repository.NewProductRepository(db)
productService:=service.NewProductService(productRepo)
productHandler:=handler.NewProductHandler(productService)

//////////////////// cart //////////////////////////////////////
cartRepo := repository.NewCartRepository(db)
cartService := service.NewCartService(cartRepo)
cartHandler := handler.NewCartHandler(cartService)

/////////////////// Wishlist ///////////////////////////////////
wishlistRepo:=repository.NewWishlistRepository(db)
wishlistService:=service.NewWishlistService(wishlistRepo)
wishlistHandler:=handler.NewWishlistHandler(wishlistService)

////////////////////// order //////////////////////////////////
orderRepo:=repository.NewOrderRepository(db)
orderService:=service.NewOrderService(orderRepo,cartRepo)
orderHandler:=handler.NewOrderHandler(orderService)
r:=gin.Default()
r.Static("/uploads","./uploads")
routes.RegisterRoutes(r,userHandler,productHandler,cartHandler,wishlistHandler,orderHandler)
r.Run(":8080")
}