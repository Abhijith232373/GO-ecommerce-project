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

r:=gin.Default()
routes.RegisterRoutes(r,userHandler,productHandler)
r.Run(":8080")
}