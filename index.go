package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
	"github.com/ssi-anik/gorm-test/connection"
	"github.com/ssi-anik/gorm-test/models"
	"fmt"
)

func main() {
	db := connection.GetConnection()
	defer db.Close()

	db.Create(&models.Product{Code: "1234", Price: 1010, Name: "Product name " + time.StampMilli})

	var product models.Product
	var products []models.Product

	db.First(&product, 1)
	db.Where("code = ?", "L1212").Find(&products)
	fmt.Println(products)
}
