package model

import (
	"gorm.io/gorm"
)

type Sanpham struct {
	gorm.Model
	MASP   string `gorm:"primarykey"`
	TENSP  string
	DVT    string
	GIA    float64
	NUOCSX string
	CTHD   []CTHD `gorm:"foreignKey:MASP;"`
}

func AddSP(db *gorm.DB, sanpham Sanpham) *Sanpham {
	db.Create(&sanpham)
	return &sanpham
}

func updateSP(db *gorm.DB, sanpham Sanpham) *Sanpham {
	db.Save(&sanpham)
	return &sanpham
}
