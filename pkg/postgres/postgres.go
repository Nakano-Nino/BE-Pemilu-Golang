package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error

	connStr := "host=localhost user=postgres password=nakanonin0 dbname=golangPractice port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database")
}
