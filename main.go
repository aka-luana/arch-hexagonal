package main

import (
	"database/sql"
	db2 "github.com/aka-luana/arch-hexagonal/adapters/db"
	"github.com/aka-luana/arch-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Produto exemplo 2", 15)

	productService.Enable(product)
}
