package database

import (
	"fmt"

	"github.com/markusremplbauer/jwtauth/backend/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(user, password, host, dbname, port string) {
	fmt.Println(user, password)
	dsn := fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%s sslmode=disable", user, password, host, dbname, port)
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	DB = connection

	if err := connection.AutoMigrate(&models.User{}); err != nil {
		return
	}
}
