package model

import (
	"gorm.io/gorm"
)

type Sanpham struct {
	gorm.Model
	MASP  string `gorm:"primarykey"`
	TENSP string
	DVT   string
	GIA   int64
	NUOC  string
	CTHD  []CTHD `gorm:"foreignKey:MASP;"`
}
