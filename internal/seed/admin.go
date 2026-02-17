package seed

import (
	"e-commerce/internal/config"
	"e-commerce/internal/models"
	"e-commerce/utils"
	"log"
)

func SeedAdmin() {
	db := config.DB
	email:="admin@gmail.com"

	var existing models.Users

	err:=db.Where("email = ?").First(&existing).Error

	if err==nil {
		log.Println("Admin already exists")
		return
	}
	hashPassword,err:=utils.HashPassword("admin123")

	if err!=nil{
		log.Fatal("Password hash failed:",err)
	}
	admin:=models.Users{
		Name: "Admin",
		Email: email,
		Password: hashPassword,
		Role: "admin",
	}
	err=db.Create(&admin).Error

	if err!=nil{
		log.Fatal("ADmin creation failed:",err)
	}
	log.Println("Admin created successfully...")
}