package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	url := "postgres://your_username:your_password@database:5432/your_database_name"
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("DBConnect!")
	return db
}


