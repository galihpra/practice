package model

import (
	"time"

	"gorm.io/gorm"
)

type PembelianModel struct {
	No_invoice       string `gorm:"type:varchar(100);primaryKey"`
	UserID           string `gorm:"type:varchar(55);foreignKey:Username"`
	CustomerID       string `gorm:"type:varchar(13);foreignKey:Hp"`
	Total            int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt         `gorm:"index"`
	DetailPembelians []DetailPembelianModel `gorm:"foreignKey:PembelianID"`
}
