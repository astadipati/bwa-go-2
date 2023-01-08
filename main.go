package main

import (
	"bwastartup/user"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	dsn := "root:Kul0nuwun@tcp(127.0.0.1:3306)/bwa_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	// panggil instance koneksi
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}
	userInput.Name = "tes from service"
	userInput.Email = "tes@gmail.com"
	userInput.Occupation = "service"
	userInput.Password = "password"

	userService.RegisterUser(userInput)

}
