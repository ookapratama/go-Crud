package controllers

import (
	"fmt"
	"modul/config"
	"modul/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// var mahasiswa models.Mahasiswa

func AllData(ctx *gin.Context) {
	var data []models.Mahasiswa
	fmt.Println(data)
	config.DB.Find(&data)
	ctx.JSON(http.StatusOK, data)
}

func Buat(ctx *gin.Context) {
	var mahasiswa models.Mahasiswa
	error := ctx.BindJSON(&mahasiswa)
	if error != nil {
		panic(error)
	}
	// fmt.Print(mahasiswa)
	config.DB.Create(&mahasiswa)
	ctx.JSON(http.StatusOK, mahasiswa)

}

func Show(ctx *gin.Context) {

	var mahasiswa models.Mahasiswa
	id := ctx.Params.ByName("id")
	config.DB.First(&mahasiswa, id)
	ctx.JSON(http.StatusOK, mahasiswa)

}

func Update(ctx *gin.Context) {

	id := ctx.Params.ByName("id")
	var mahasiswa models.Mahasiswa
	error := config.DB.First(&mahasiswa, id).Error
	if error != nil {
		panic(error)
	}

	if error := ctx.BindJSON(&mahasiswa); error != nil {
		panic(error)
	}

	config.DB.Updates(&mahasiswa)
	ctx.JSON(http.StatusOK, mahasiswa)

}

func Delete(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	
	var mahasiswa models.Mahasiswa
	error := config.DB.First(&mahasiswa, id).Error
	if error != nil {
		panic(error)
	}
	config.DB.Delete(&mahasiswa)
	ctx.JSON(http.StatusOK, mahasiswa)

}
