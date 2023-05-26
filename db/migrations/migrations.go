package main

import (
	"log"
	"user-go/db"
	"user-go/models"
)

func main() {
	db := db.ConnectionDB()
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

}
