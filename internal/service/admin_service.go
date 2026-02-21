package service

import (
	"e-commerce/internal/models"
	"e-commerce/internal/repository"
	"e-commerce/utils"
	"fmt"

	// "golang.org/x/text/search"
)

type AdminService struct {
	adminRepo *repository.AdminRepository
}

func  NewAdminService(adminRepo *repository.AdminRepository) *AdminService{
	return &AdminService{
		adminRepo: adminRepo,
	}
}


func (s *AdminService)Login(email,password string)(*models.Users,error){
	user,err:=s.adminRepo.FindByEmail(email)
	if err !=nil{
		return nil,err
	}
	if user.Role !="admin"{
		return nil,fmt.Errorf("not admin")
	}
	err = utils.CheckPassword(user.Password, password)
	
	if err !=nil{
		return nil,err
	}
	return user,nil
}

/////////////// Users ///////////////////////
func (s *AdminService)GetAllUsers()([]models.Users,error){
	return s.adminRepo.GetAllUsers()
}

func (s *AdminService) AddUser(name, email, password string) error {
	return s.adminRepo.AddUser(name, email, password)
}

func  (s *AdminService)SuspendUser(id string)error{
	return s.adminRepo.SuspendUser(id)
}

func (s *AdminService)ActiveUser(id string)error{
	return s.adminRepo.ActivateUser(id)
}

func (s *AdminService)DeleteUser(id string)error{
	return s.adminRepo.DeleteUser(id)
}


////////////////////  Products ///////////////////////
func (s *AdminService)GetAllProducts()([]models.Product,error){
	return s.adminRepo.GetAllProducts()
}
func (s *AdminService)AddProduct(name string,description string, price float64,stock int,categoryID uint,image string)error{
	return  s.adminRepo.AddProduct(name,description,price,stock,categoryID,image)
}
func (s *AdminService)DeleteProduct(id string)error{
	return s.adminRepo.DeleteProduct(id)
}

func (s *ProductService)ActivateProduct(id string)error{
	return s.ProductRepo.ActivateProduct(id)
}

func (s  *ProductService)SuspendProduct(id string)error{
	return s.ProductRepo.SuspendProduct(id)
}

func  (s *AdminService)SearchProducts(search string)([]models.Product,error){
	return s.adminRepo.SearchProducts(search)
}