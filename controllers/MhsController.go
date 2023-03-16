package controllers

import (
	"modul/config"
	"modul/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status bool
	Pesan string
	data interface{}
}

func AllData(ctx *gin.Context) {
	var data []models.Mahasiswa
	config.DB.Find(&data)

	var Res Response
	Res.Status = true
	Res.Pesan = "Data Mahasiswa"
	Res.data = data

	ctx.JSON(http.StatusOK, Res)
}