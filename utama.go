package main

import (
	"log"
	"modul/config"
	"modul/controllers"
	"modul/models"
	"net/http"
	"os"

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
	rute.POST("/tambah", controllers.Buat)
	rute.GET("/show/:id", controllers.Show)
	rute.PUT("/update/:id", controllers.Update)
	rute.DELETE("/destroy/:id", controllers.Delete)

	rute.Run(":" + os.Getenv("APP_PORT"))

}
