package main

import (
	"log"

	"github.com/what-crud/initializers"
	"github.com/what-crud/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Migrate successfully!")
}
