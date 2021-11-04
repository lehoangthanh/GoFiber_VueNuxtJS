package main

import (
	"github.com/bxcodec/faker/v3"
	"hello/src/database"
	"hello/src/models"
)

func main() {
	database.Connect()
	for i :=0; i< 30; i++ {
		ambassador := models.User{
			FirstName: faker.FirstName(),
			LastName: faker.LastName(),
			Email: faker.Email(),
			IsAmbassador: true,
		}
		ambassador.SetPassword("123")
		database.DB.Create(&ambassador)
	}
}