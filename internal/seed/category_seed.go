package seed

import (
	"e-commerce/internal/config"
	"e-commerce/internal/models"
	"log"
)

func SeedCategories() {
	db := config.DB

	categories:=[]models.Category{
		{Name: "Chairs"},
		{Name: "Sofas"},
		{Name: "Tables"},
		{Name: "Beds"},
		{Name: "Storage"},
		{Name: "Shelves & TV Units"},
	}

	for _,category:=range categories{
		var existing models.Category

		err:=db.Where("name=?",category.Name).First(&existing).Error

		if err == nil {
			 log.Println("Category exists:",category.Name)
			 continue
		}
		db.Create(&category)

		log.Println("Inserted category:",category.Name)
	}
}