package routes
import (
	"e-commerce/controllers"
	"github.com/gin-gonic/gin"
)
func AuthRoutes(r *gin.Engine){
	auth:=r.Group("/api/auth")
	{
		auth.POST("/register",controllers.Register)
		auth.POST("/login",controllers.Login)
		auth.POST("/logout",controllers.Logout)
	}
}