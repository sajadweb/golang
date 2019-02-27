package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open("sqlite3", "./test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "ddd", Price: 1000})

	// Read
	var product Product
	db.First(&product, 2)                  // find product with id 1
	db.First(&product, "code = ?", "dddd") // find product with code l1212

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 8000)

	// Delete - delete product
	db.Delete(&product)

}
