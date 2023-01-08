package main

import (
	"bwastartup/user"
	"fmt"
	"log"
	"os"

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
	// username := os.Getenv("USERE")
	// fmt.Println(username)
	// pass := os.Getenv("PASS")
	// fmt.Println(pass)
	dsn := "root:Kul0nuwun@tcp(127.0.0.1:3306)/bwa_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Koneksi Ok")

	var users []user.User
	length := len(users)
	fmt.Println(length)

	db.Find(&users)

	length = len(users)
	fmt.Println(length)
	fmt.Printf("%s hh %s\n", os.Getenv("USER"), os.Getenv("PASSWORD"))

	for _, user := range users {
		fmt.Println(user.Name)
		fmt.Println(user.Email)
	}
}
