package main

import (
	"github.com/bxcodec/faker/v3"
	"hello/src/database"
	"hello/src/models"
	"math/rand"
)

func main() {
	database.Connect()
	for i := 0; i < 30; i++ {
		ambassador := models.Product{
			Title:       faker.FirstName(),
			Description: faker.LastName(),
			Image:       faker.URL(),
			Price:       float64(rand.Intn(90) + 10),
		}
		database.DB.Create(&ambassador)
	}
}