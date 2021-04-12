package main

import (
	"log"
	"net/http"

	"github.com/fernandocastrotelco/gobackend/inventory/inventory-service/database"
	"github.com/fernandocastrotelco/gobackend/inventory/inventory-service/product"
	"github.com/fernandocastrotelco/gobackend/inventory/inventory-service/receipt"
	_ "github.com/go-sql-driver/mysql"
)

const basePath = "/api"

func main() {
	database.SetupDatabase()
	receipt.SetupRoutes(basePath)
	product.SetupRoutes(basePath)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
