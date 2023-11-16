package model

import (
	"time"

	"gorm.io/gorm"
)

type ProductModel struct {
	Barcode   string `gorm:"type:varchar(20);primaryKey"`
	UserID    string
	Nama      string
	Harga     int
	Stok      int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserNama  string         `gorm:"column:nama"`
}
