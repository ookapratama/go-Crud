package models

import (
	"time"

	"gorm.io/gorm"
)

type Mahasiswa struct {
	ID        uint64         `gorm:"primary_key:auto_increment" json:"id"`
	Nama      string         `json:"nama" binding:"required"`
	Stambuk   string         `json:"stambuk" binding:"required" gorm:"size:6"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
