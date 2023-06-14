package model

import (
	"gorm.io/gorm"
	"time"
)

type Hoadon struct {
	gorm.Model
	SOHD   int `gorm:"primaryKey;autoIncrement"`
	NGHD   time.Time
	MAKH   string
	MANV   string
	TRIGIA int
}

func AddHoaDon(db *gorm.DB, hoadon Hoadon) *Hoadon {
	db.Create(&hoadon)
	return &hoadon
}
