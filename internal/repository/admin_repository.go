package repository

import (
	"e-commerce/internal/config"
	"e-commerce/internal/models"

	// "image"

	// "golang.org/x/text/search"
	"gorm.io/gorm"
)

// db *gorm.DB

type AdminRepository struct{
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{db: db}
}

func (r *AdminRepository)GetAllProducts() ([]models.Product,error){
	var products []models.Product
	err:=config.DB.Find(&products).Error
	return products,err
}

func(r *AdminRepository)AddProduct(name string,description string,price float64,stock int,categoryID uint, image string)error{
	product:=models.Product{
		Name: name,
		Description: description,
		Price: price,
		Stock: stock,
		CategoryID:categoryID,
		ImageURL: image,
		IsActive: true,
	}
	return  r.db.Create(&product).Error
}

func (r *AdminRepository)DeleteProduct(id string)error{
	return r.db.Delete(&models.Product{},id).Error
}

func (r *AdminRepository)FindByEmail(email string)(*models.Users,error){
	var user models.Users
	err :=config.DB.Where("email = ?",email).First(&user).Error
	return &user,err
}

func (r *AdminRepository)SearchProducts(search string)([]models.Product,error){
	var products []models.Product

	err :=r.db.
	Where("name ILIKE ?","%"+search+"%").
	Find(&products).Error

	return products,err
}



////////////////////// users//////////////////
func (r *AdminRepository)GetAllUsers()([]models.Users,error){
	var users []models.Users

	err :=r.db.Find(&users).Error

	return users,err
}
func (r *AdminRepository)AddUser(name, email, password string) error {

	user := models.Users{
		Name: name,
		Email: email,
		Password: password,
		IsActive: true,
	}

	return r.db.Create(&user).Error
}

func (r *AdminRepository)SuspendUser(id string)error{
	return r.db.Model(&models.Users{}).
	Where("id = ?",id).
	Update("is_active",false).Error

}
func (r *AdminRepository)ActivateUser(id string) error {
	return r.db.Model(&models.Users{}).
	Where("id = ?",id).
	Update("is_active",true).Error
}
func (r *AdminRepository)DeleteUser(id string)error{
	return r.db.Delete(&models.Users{},id).Error
}




