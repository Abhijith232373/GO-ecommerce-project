package repository

import (
	"e-commerce/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

///////////////////////// Create //////////////////////

func (r *UserRepository)Create(user *models.Users)error{
	return r.DB.Create(user).Error
}

////////////////////////// Find by Email //////////////

func (r *UserRepository)FindByEmail(email string)(*models.Users,error){
	var user models.Users
	err :=r.DB.Where("email=?",email).First(&user).Error
	if err!=nil{
		return nil,err
	}
	return  &user,nil
}

/////////////////////// GET BY ID ///////////////////////////

func (r *UserRepository)GetByID(userID uint)(*models.Users,error){
	var user models.Users
	err:=r.DB.First(&user,userID).Error
	if err !=nil{
		return nil,err
	}
	return &user,nil
}

//////////////////////// Update  ///////////////////////////

func (r *UserRepository)Update(user *models.Users)error{
	return r.DB.Save(user).Error
}