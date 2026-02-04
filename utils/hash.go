package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashed,err:=bcrypt.GenerateFromPassword([]byte(password),14)
	return string(hashed),err
}

func ComaparePassword(hash,password string)error{
	return bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
}