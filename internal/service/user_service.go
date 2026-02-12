package service

import (
	"e-commerce/internal/models"
	"e-commerce/internal/repository"
	"e-commerce/utils"
	"errors"
	"strings"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{UserRepo: repo}
}

//////////////////////////////// Register ///////////////////////////////

func (s *UserService) Register(user *models.Users) error {
	user.Email = strings.ToLower(strings.TrimSpace(user.Email))
	user.Password = strings.TrimSpace(user.Password)

	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hash
	return s.UserRepo.Create(user)
}

////////////////////////////// LOGIN  //////////////////////////////////

func (s *UserService) Login(email, password string) (string, string, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	password = strings.TrimSpace(password)

	user, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return "", "", errors.New("invalid email")
	}
	
	if err := utils.CheckPassword(user.Password, password); err != nil {
		return "", "", errors.New("invalid password")
	}
			
	accessToken, err := utils.GenerateAccessToken(user.ID, user.Email, user.Role)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		return "", "", err
	}

	user.RefreshToken=refreshToken
	if err:=s.UserRepo.Update(user);err!=nil{
		return "","",err
	}
	return accessToken, refreshToken, nil
}

///////////////////////////////  Refresh /////////////////////////

func (s *UserService)Refresh(refreshToken string)(string,error){

	claims,err:=utils.ValidateRefreshToken(refreshToken)
	if err!=nil{
		return "",err
	}
	user,err:=s.UserRepo.GetByID(claims.UserID)
	if err!=nil{
		return "",err
	}
	if user.RefreshToken !=refreshToken{
		return "",errors.New("invalid refresh token")
	}
	newAccessToken,err:=utils.GenerateAccessToken(
		user.ID,
		user.Email,
		user.Role,
	)
	if err !=nil{
		return "",err
	}
	return newAccessToken,nil
}

//////////////////////////// Logout ////////////////////////

func (s *UserService)Logout(userID uint)error{
	user,err:=s.UserRepo.GetByID(userID)
	if err != nil{
		return err
	}
	user.RefreshToken = ""
	
	return s.UserRepo.Update(user)
}