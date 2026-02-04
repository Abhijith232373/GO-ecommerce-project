package config

import "github.com/joho/godotenv"

func Loadenv() {
	err:=godotenv.Load()
	if err!=nil{
		panic("Error loading from .env file")
	}
}