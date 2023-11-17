package model

import (
	"time"

	"gorm.io/gorm"
)

type DetailPembelianModel struct {
	PembelianID string `gorm:"type:varchar(100);primaryKey"`
	ProductID   string `gorm:"type:varchar(20);primaryKey"`
	Sub_total   int
	Qty         int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	ProdukNama  string         `gorm:"column:nama"`
}

type DetailPembelianQuery struct {
	DB *gorm.DB
}

func (dpq *DetailPembelianQuery) AddDetailPembelian(newDetailPembelian DetailPembelianModel) (DetailPembelianModel, error) {
	if err := dpq.DB.Create(&newDetailPembelian).Error; err != nil {
		return DetailPembelianModel{}, err
	}

	return newDetailPembelian, nil
}

func (dpq *DetailPembelianQuery) ListDetailPembelian() ([]DetailPembelianModel, error) {
	var detailPembelians []DetailPembelianModel
	if err := dpq.DB.Find(&detailPembelians).Error; err != nil {
		return nil, err
	}

	return detailPembelians, nil
}

func (dpq *DetailPembelianQuery) GetDetailPembelianByInvoice(no_invoice string) (DetailPembelianModel, error) {
	var detailPembelians DetailPembelianModel
	if err := dpq.DB.Where("no_invoice = ?", no_invoice).First(&detailPembelians).Error; err != nil {
		return DetailPembelianModel{}, err
	}

	return detailPembelians, nil
}

func (dpq *DetailPembelianQuery) UpdateDetailPembelian(no_invoice string, updatedData DetailPembelianModel) (DetailPembelianModel, error) {
	var detailPembelians DetailPembelianModel
	if err := dpq.DB.Model(&detailPembelians).Where("no_invoice = ?", no_invoice).Updates(updatedData).Error; err != nil {
		return DetailPembelianModel{}, err
	}

	return detailPembelians, nil
}

func (dpq *DetailPembelianQuery) DeleteDetailPembelian(no_invoice string) error {
	var detailPembelians DetailPembelianModel
	if err := dpq.DB.Model(&detailPembelians).Where("no_invoice = ?", no_invoice).Delete(&detailPembelians).Error; err != nil {
		return err
	}

	return nil
}
