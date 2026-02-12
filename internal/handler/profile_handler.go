package handler

import (

	"github.com/gin-gonic/gin"
)

func (h *UserHandler)Profile(c *gin.Context){
	userID:=c.GetUint("userID")

	user,err:=h.UserService.UserRepo.GetByID(userID)

	if err!=nil{
		c.JSON(400,gin.H{"error":"user not found !"})
		return	
	}
	c.JSON(200,gin.H{
		"id":user.ID,
		"name":user.Name,
		"email":user.Email,
		"role":user.Role,
	})
}