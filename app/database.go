package app

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func DBConnection() *gorm.DB {

	dsn := "user=postgres password=123 dbname=test port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	return db

}
