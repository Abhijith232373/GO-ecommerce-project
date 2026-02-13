package service

import (
	"e-commerce/internal/models"
	"e-commerce/internal/repository"
)

type WishlistService struct {
	Repo *repository.WishlistRepository
}

func NewWishlistService(repo *repository.WishlistRepository)*WishlistService{
	return &WishlistService{Repo: repo}
}
func (s *WishlistService)Add(userID,productID uint)error{
	w:=models.Wishlist{
		UserID: userID,
		ProductID: productID,
	}
	return s.Repo.Add(&w)
}

func (s *WishlistService)Get(userID uint)([]models.Wishlist,error){
	return s.Repo.GetByUser(userID)
}

func (s *WishlistService)Remove(id uint)error{
	return s.Repo.Remove(id)
}