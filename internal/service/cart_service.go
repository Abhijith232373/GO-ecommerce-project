package service

import (
	"e-commerce/internal/models"
	"e-commerce/internal/repository"
	"errors"
)

type CartService struct {
	CartRepo *repository.CartRepository
}

func NewCartService(repo * repository.CartRepository)*CartService{
	return &CartService{CartRepo: repo}
}

func (s *CartService)AddToCart(userID uint,ProductID uint,qty int)error {
	cart:=models.Cart{
		UserID: userID,
		ProductID: ProductID,
		Quantity: qty,
	}
	return  s.CartRepo.Add(&cart)
}
func (s *CartService)GetCart(userID uint)([]models.Cart,error){
	return  s.CartRepo.GetByUser(userID)
}
func (s *CartService)Remove(cartID uint)error{
	return s.CartRepo.Delete(cartID)
}

func (s *CartService)UpdateQuantity(cartID uint,quantity int)error{
	if quantity < 1 {
		return errors.New("quantity must be at least 1")
	}
	return s.CartRepo.UpdateQuantity(cartID,quantity)
}