package connection

import (
	"os"
	"fmt"
	"github.com/jinzhu/gorm"
)

func GetConnection() (db *gorm.DB) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_DATABASE")
	password := os.Getenv("DB_PASSWORD")

	connection := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbName, password)

	db, err := gorm.Open("postgres", connection)

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
