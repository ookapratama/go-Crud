package main

import (
	"log"
	"modul/config"
	"modul/models"

	"github.com/joho/godotenv"
)

func init() {
	error := godotenv.Load()
	if error != nil {
		log.Fatal("Error saat membaca file .env")
	}

	config.ConnDB()
	config.DB.AutoMigrate(&models.Mahasiswa{})

}

func main() {
	// fmt.Println("Fungsi Utama")
}
