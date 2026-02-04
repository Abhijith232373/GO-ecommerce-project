package main

import (
	"e-commerce/config"
	"e-commerce/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)
func main() {
	config.Loadenv()
	config.ConnectDB()

	r:=gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowOrigins:[]string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
		AllowHeaders: []string{"Origin","Content-Type","Authorization"},
		AllowCredentials: true,
		MaxAge: 12*60*60,
	}))
	r.OPTIONS("/*path",func(c *gin.Context) {
		c.Status(204)
	})
	routes.AuthRoutes(r)

	r.GET("/health",func(c *gin.Context) {
		c.JSON(200,gin.H{"status":"ok"})
	})
	r.Run(":8080")
}