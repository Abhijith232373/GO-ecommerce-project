package seed

import "log"

func SeedAll() {
	// log.Println("Starting databse seeding...")
	log.Println("Starting databse seeding...")
	SeedCategories()
	SeedProducts()  
	
	SeedAdmin()
	log.Println("Seed finished successfully")

}