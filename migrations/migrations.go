package main

import (
	"github.com/ssi-anik/gorm-test/models"
	"github.com/ssi-anik/gorm-test"
)

func Migrations() {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Product{})
	defer db.Close()
}

func main() {
	Migrations()
}
