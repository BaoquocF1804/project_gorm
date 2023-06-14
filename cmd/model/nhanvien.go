package model

import (
	"gorm.io/gorm"
	"time"
)

type Nhanvien struct {
	gorm.Model
	MANV   string `gorm:"primarykey"`
	HOTEN  string
	NGVL   time.Time
	SODT   string
	Hoadon []Hoadon `gorm:"foreignKey:MANV"`
}

func AddNhanvien(db *gorm.DB, nhanvien Nhanvien) *Nhanvien {
	db.Create(&nhanvien)
	return &nhanvien
}

func DeleteNV(db *gorm.DB, MANV string) {
	db.Delete(&Nhanvien{}, "MANV = ?", MANV)
	return
}
