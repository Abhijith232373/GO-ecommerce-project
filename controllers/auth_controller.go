package controllers

import (
	"e-commerce/config"
	"e-commerce/models"
	"e-commerce/utils"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

// //////////////////////////////////////      signup     /////////////////////////////////////

func Register(c *gin.Context){
	var input struct{
		Name string `json:"name"`
		Email string `json:"email"`
		Password string `json:"password"`
	}
	fmt.Println(" register Api  hit")
	if err:=c.ShouldBindJSON(&input);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	fmt.Println("payload",input)
	hashed,_:=utils.HashPassword(input.Password)
	user:=models.User{
		Name: input.Name,
		Email: input.Email,
		Password: hashed,
		Role: "user",
	}
	if err:=config.DB.Create(&user).Error;err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"email already exists"})
		return
	}
	c.JSON(http.StatusCreated,gin.H{
		"message":"user regitered successfully !",
	})
}


////////////////////////////////////////      Login     /////////////////////////////////////


func Login(c *gin.Context){
	var  input struct{
		Email  string  `json:"email"`
		Password string `json:"password"`
	}

	var user models.User
	if err:=c.ShouldBindJSON(&input);err!=nil{
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}

	if err:=config.DB.Where("email=?",input.Email).First(&user).Error;err!=nil{
		c.JSON(401,gin.H{"error":"invalid cendentials !"})
		return
	}

	if err:=utils.ComaparePassword(user.Password,input.Password);err!=nil{
		c.JSON(401,gin.H{"error":"invalid cendertails !"})
		return
	}
	
	token,_:=utils.GenernateToken(user.ID,user.Role)
	c.JSON(200,gin.H{
		"token":token,
		"user":gin.H{
			"id":user.ID,
			"name":user.Name,
			"email":user.Email,
			"role":user.Role,
		},
	})
}


func Logout(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"message":"logout successfully",
	})
}