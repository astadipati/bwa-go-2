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
	// panggil struct
	user := user.User{
		Name:           "doni",
		Occupation:     "programmer",
		Email:          "doni@gmail.com",
		PasswordHash:   "123456",
		AvatarFileName: "doni.jpg",
		Role:           "admin",
		// CreatedAt: "2023-01-01",
	}
	// punya fungsi save
	userRepository.Save(user)

}
