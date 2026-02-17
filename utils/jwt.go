package utils

import (
	"errors"
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

func getAccessSecret() []byte {
	return []byte(os.Getenv("JWT_ACCESS_SECRET"))
}

func getRefreshSecret() []byte {
	return []byte(os.Getenv("JWT_REFRESH_SECRET"))
}

///////////////////////////////  AccessToken //////////////////////////

type AccessClaims struct {
	UserId uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	// TokenVersion int `json:"token_version"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(userId uint, email, role string)(string, error) {
	claims := AccessClaims{
		UserId: userId,
		Email:  email,
		Role:   role,
		// TokenVersion: tokenVersion,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getAccessSecret())
}

//////////////////////////////////// RefreshToken ////////////////////////////////

type RefreshClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateRefreshToken(userID uint) (string, error) {
	claims := RefreshClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "ecommerce-app",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getRefreshSecret())
}

/////////////////////////////// Validation /////////////////////////////////// 


func ValidateAccessToken(tokenStr string) (*AccessClaims, error) {
	claims := &AccessClaims{}
	token,err:=jwt.ParseWithClaims(
		tokenStr,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			if token.Method !=jwt.SigningMethodHS256{
				return nil,errors.New("invalid signing method")
			}
			return getAccessSecret(),nil
		},
	)
	if err!=nil || !token.Valid{
		return  nil,errors.New("invalid or expire token !")
	}
	return claims,nil
}

func ValidateRefreshToken(tokenStr string) (*RefreshClaims, error) {
	claims := &RefreshClaims{}
	token,err:=jwt.ParseWithClaims(
		tokenStr,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			if _,ok:=token.Method.(*jwt.SigningMethodHMAC);!ok{
				return nil,errors.New("invalid signing method")
			}
			return  getRefreshSecret(),nil
		},
	)
	if err!=nil || !token.Valid {
		return nil,errors.New("invalid or expired refresh token")
	}
	return claims,nil
}

