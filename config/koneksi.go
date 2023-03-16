package config

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnDB() {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	usr := os.Getenv("DB_USER")
	psw := os.Getenv("DB_PASSWORD")
	db := os.Getenv("DB_NAME")

	dsn := usr + ":" + psw + "@(" + host + ":" + port + ")/" + db + "?charset=utf8mb4&parseTime=True&loc=Local" 

	config := gorm.Config{}

	database, error := gorm.Open(mysql.Open(dsn), &config)
	if error != nil {
		panic("Gagal Terhubung ke Database " + os.Getenv("DB_NAME"))
	}

	DB = database

}

func CloseConnDB(database *gorm.DB) {
	dbSQL, error := database.DB()
	if error != nil {
		panic("Gagal koneksi ke Database")
	}
	dbSQL.Close()
}






