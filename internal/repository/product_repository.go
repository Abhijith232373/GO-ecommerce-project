package repository

import (
	"e-commerce/internal/models"
	"gorm.io/gorm"
)

type ProductRepository struct{
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository{
	return &ProductRepository{DB:db}
}

func (r *ProductRepository)GetAll(
	search string,
	category string,
	page int,
	limit int,
	) ([]models.Product,int64,error){
	var products []models.Product
	var total int64

	query:=r.DB.Model(&models.Product{}).
	Where("stock>?",0)

	if search !=""{
		// searchPattern:="%"+strings.ToLower(search)+"%"
		query=query.Where("LOWER(name)LIKE LOWER(?)","%"+search+"%")
	}
	query.Count(&total)
	offset:=(page -1) * limit

	err:=query.
		Preload("Category").
		Limit(limit).
		Offset(offset).
		Find(&products).Error

	return  products,total,err
}

func (r *ProductRepository) ActivateProduct(id string) error {
	return r.DB.Model(&models.Product{}).
		Where("id = ?", id).
		Update("is_active", true).Error
}

func (r *ProductRepository) SuspendProduct(id string) error {
	return r.DB.Model(&models.Product{}).
		Where("id = ?", id).
		Update("is_active", false).Error
}
