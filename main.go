package main

import (

	// mysqlDb "github.com/Sahil-Mandaliya/Order/backend/pkg/mysql"
	// routes "github.com/Sahil-Mandaliya/Order/backend/routes"

	"log"

	routes "github.com/Sahil-Mandaliya/OrderService/routes"

	mysqlDb "github.com/Sahil-Mandaliya/OrderService/pkg/mysql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := mysqlDb.InitDB()
	if db.Error != nil {
		log.Fatalln(db.Error)
	}
	connection := mysqlDb.DBConnection()
	condb := connection.DB()
	defer condb.Close()
	routes.Routers()
}
