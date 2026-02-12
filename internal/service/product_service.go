package service

import (
	"e-commerce/internal/models"
	"e-commerce/internal/repository"
)
type ProductService struct{
	ProductRepo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{ProductRepo: repo}
}

func (s *ProductService)GetProducts(
	search string,
	category string,
	page int,
	limit int,
	)([]models.Product,int64,error){
	
	return  s.ProductRepo.GetAll(search,category,page,limit)
}

