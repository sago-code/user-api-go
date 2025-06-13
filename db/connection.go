package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DSN = "root:@tcp(localhost:3306)/user-api-go?charset=utf8mb4&parseTime=True&loc=Local"
var DB *gorm.DB

func DBConnection() {
	var err error
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	} else {
		log.Println("DB connected")
	}
}
