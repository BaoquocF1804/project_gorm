package model

import (
	"gorm.io/gorm"
	"time"
)

type Khachhang struct {
	gorm.Model
	MAKH    string `gorm:"primarykey"`
	HOTEN   string
	DCHI    string
	SODT    string
	NGSINH  time.Time
	DOANHSO float64
	NGDK    time.Time
	Hoadon  []Hoadon `gorm:"foreignKey:MAKH;"`
}
