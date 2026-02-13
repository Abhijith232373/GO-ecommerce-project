package repository

import (
	"e-commerce/internal/models"
	"gorm.io/gorm"
)

type WishlistRepository struct {
	DB *gorm.DB
}

func NewWishlistRepository(db *gorm.DB)*WishlistRepository{
	return &WishlistRepository{DB: db}
}

func (r *WishlistRepository)Add(w *models.Wishlist)error{
	return r.DB.Create(w).Error
}

func (r *WishlistRepository)GetByUser(userID uint)([]models.Wishlist,error){
	var data []models.Wishlist

	err:=r.DB.
	Preload("Product").
	Where("user_id=?",userID).
	Find(&data).Error

	return data,err
}

func (r *WishlistRepository)Remove(id uint)error{
	return r.DB.Delete(&models.Wishlist{},id).Error
}