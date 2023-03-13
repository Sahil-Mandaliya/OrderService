package mysql

// import (
// 	"log"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// func InitDB() error {
// 	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
// 	dsn := "root:@tcp(127.0.0.1:3306)/userdb?charset=utf8mb4&parseTime=True&loc=Local"
// 	gdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return err
// 	}
// 	db = gdb
// 	log.Println("Database initialised successfully")
// 	return nil
// }

import (
	"fmt"
	"log"

	config "github.com/Sahil-Mandaliya/OrderService/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error
var DB *gorm.DB

func getMySQLConnectionString() string {
	DBUser := config.DBUser()
	DBPassword := config.DBPassword()
	DBHost := config.DBHost()
	DBPort := config.DBPort()
	DBName := config.DBName()
	dataBase := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DBUser,
		DBPassword,
		DBHost,
		DBPort,
		DBName)

	// dataBase = "user:password@tcp(localhost:3308)/order?charset=utf8mb4&parseTime=True&loc=Local"
	return dataBase
}

func InitDB(params ...string) *gorm.DB {
	var err error
	conString := getMySQLConnectionString()

	DB, err = gorm.Open("mysql", conString)

	if err != nil {
		log.Panic(err)
	}
	db = DB
	log.Println("Database initialised successfully")
	return DB
}

func DBConnection() *gorm.DB {
	return db
}
