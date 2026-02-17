package seed

import (
	"log"

	"e-commerce/internal/config"
	"e-commerce/internal/models"
)

func SeedProducts() {

	db := config.DB

	if db == nil {
		log.Fatal("Database not initialized")
	}

	products := []models.Product{

		{
			Name:        "Wooden Dining Chair",
			Description: "Solid wood dining chair",
			Price:       2500,
			Stock:       20,
			CategoryID:  1,
			ImageURL:    "chair1.jpg",
		},

		{
			Name:        "Office Chair",
			Description: "Ergonomic office chair",
			Price:       5500,
			Stock:       15,
			CategoryID:  1,
			ImageURL:    "chair2.jpg",
		},

		{
			Name:        "Luxury Sofa",
			Description: "Premium 3 seater sofa",
			Price:       25000,
			Stock:       6,
			CategoryID:  2,
			ImageURL:    "sofa1.jpg",
		},

		{
			Name:        "L Shape Sofa",
			Description: "Modern L shape sofa",
			Price:       38000,
			Stock:       4,
			CategoryID:  2,
			ImageURL:    "sofa2.jpg",
		},

		{
			Name:        "Coffee Table",
			Description: "Wooden coffee table",
			Price:       4500,
			Stock:       12,
			CategoryID:  3,
			ImageURL:    "table1.jpg",
		},

		{
			Name:        "Dining Table 6 Seater",
			Description: "Family dining table",
			Price:       18000,
			Stock:       7,
			CategoryID:  3,
			ImageURL:    "table2.jpg",
		},

		{
			Name:        "King Size Bed",
			Description: "King size storage bed",
			Price:       32000,
			Stock:       5,
			CategoryID:  4,
			ImageURL:    "bed1.jpg",
		},

		{
			Name:        "Queen Size Bed",
			Description: "Queen size wooden bed",
			Price:       24000,
			Stock:       8,
			CategoryID:  4,
			ImageURL:    "bed2.jpg",
		},

		{
			Name:        "Wooden Wardrobe",
			Description: "3 door wardrobe",
			Price:       21000,
			Stock:       6,
			CategoryID:  5,
			ImageURL:    "wardrobe1.jpg",
		},

		{
			Name:        "Sliding Wardrobe",
			Description: "Modern sliding wardrobe",
			Price:       35000,
			Stock:       3,
			CategoryID:  5,
			ImageURL:    "wardrobe2.jpg",
		},

		{
			Name:        "Bookshelf",
			Description: "5 shelf wooden bookshelf",
			Price:       6000,
			Stock:       10,
			CategoryID:  6,
			ImageURL:    "bookshelf1.jpg",
		},

		{
			Name:        "TV Unit",
			Description: "Modern TV cabinet",
			Price:       12000,
			Stock:       9,
			CategoryID:  6,
			ImageURL:    "tvunit1.jpg",
		},

		{
			Name:        "Office Desk",
			Description: "Modern office desk",
			Price:       7500,
			Stock:       14,
			CategoryID:  3,
			ImageURL:    "desk1.jpg",
		},

		{
			Name:        "Dressing Table",
			Description: "Wood dressing table with mirror",
			Price:       9500,
			Stock:       7,
			CategoryID:  5,
			ImageURL:    "dress1.jpg",
		},

		{
			Name:        "Bedside Table",
			Description: "Nightstand bedside table",
			Price:       3200,
			Stock:       18,
			CategoryID:  5,
			ImageURL:    "night1.jpg",
		},
	}

	for _, product := range products {

		var existing models.Product

		err := db.Where("name = ?", product.Name).First(&existing).Error

		if err == nil {
			log.Println("Already exists:", product.Name)
			continue
		}

		err = db.Create(&product).Error

		if err != nil {
			log.Println("Insert failed:", product.Name)
			continue
		}

		log.Println("Inserted:", product.Name)
	}

	log.Println("products seed completed...")
}