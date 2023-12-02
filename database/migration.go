package database

import (
	"Restful_Go/models"
	"Restful_Go/pkg/postgres"
	"fmt"
)

func RunMigration() {
	err := postgres.DB.AutoMigrate(&models.User{}, &models.Paslon{}, &models.Article{}, &models.Partai{}, &models.Voter{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Migration Success")
}
