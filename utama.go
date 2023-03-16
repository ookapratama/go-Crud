package main

import (
	"log"
	"modul/config"
	"modul/controllers"
	"modul/models"
	"net/http"

	"github.com/gin-gonic/gin"
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

	rute := gin.Default()

	rute.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"pesan": "pong",
		})
	})

	rute.GET("/index", controllers.AllData)

	rute.Run()

}
