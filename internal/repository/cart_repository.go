package repository

import (
	"e-commerce/internal/models"

	"gorm.io/gorm"
)

type CartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB)*CartRepository{
	return  &CartRepository{DB: db}
}

func (r *CartRepository)Add(cart *models.Cart)error{
	return r.DB.Create(cart).Error
}

func (r *CartRepository)GetByUser(userID uint)([]models.Cart,error){
	var carts []models.Cart

	err:=r.DB.
	Preload("Product").
	Where("user_id=?",userID).
	Find(&carts).Error

	return carts,err
}

func (r *CartRepository)Delete(cartID uint)error{
	return r.DB.Delete(&models.Cart{},cartID).Error
}

func (r *CartRepository)UpdateQuantity(cartID uint,quantity int)error {
	return r.DB.Model(&models.Cart{}).
	Where("id=?",cartID).
	Update("quantity",quantity).Error
}