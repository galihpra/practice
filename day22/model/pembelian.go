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

type PembelianQuery struct {
	DB *gorm.DB
}

func (pbq *PembelianQuery) AddPembelian(newPembelian PembelianModel) (PembelianModel, error) {
	if err := pbq.DB.Create(&newPembelian).Error; err != nil {
		return PembelianModel{}, err
	}

	return newPembelian, nil
}

func (pbq *PembelianQuery) ListPembelian() ([]PembelianModel, error) {
	var pembelians []PembelianModel
	if err := pbq.DB.Find(&pembelians).Error; err != nil {
		return nil, err
	}

	return pembelians, nil
}

func (pbq *PembelianQuery) GetPembelianByInvoice(invoice string) (PembelianModel, error) {
	var pembelian PembelianModel
	if err := pbq.DB.Where("invoice = ?", invoice).First(&pembelian).Error; err != nil {
		return PembelianModel{}, err
	}

	return pembelian, nil
}

func (pbq *PembelianQuery) UpdatePembelian(invoice string, updatedData PembelianModel) (PembelianModel, error) {
	var pembelian PembelianModel
	if err := pbq.DB.Model(&pembelian).Where("invoice = ?", invoice).Updates(updatedData).Error; err != nil {
		return PembelianModel{}, err
	}

	return pembelian, nil
}

func (pbq *PembelianQuery) DeletePembelian(invoice string) error {
	var pembelian PembelianModel
	if err := pbq.DB.Model(&pembelian).Where("invoice = ?", invoice).Delete(&pembelian).Error; err != nil {
		return err
	}

	return nil
}
