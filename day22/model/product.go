package model

import (
	"time"

	"gorm.io/gorm"
)

type ProductModel struct {
	Barcode          string `gorm:"type:varchar(20);primaryKey"`
	UserID           string
	Nama             string
	Harga            int
	Stok             int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt         `gorm:"index"`
	UserNama         string                 `gorm:"column:user_nama"`
	DetailPembelians []DetailPembelianModel `gorm:"foreignKey:PembelianID"`
}

type ProductQuery struct {
	DB *gorm.DB
}

func (pq *ProductQuery) AddProduct(newProduct ProductModel) (ProductModel, error) {
	if err := pq.DB.Create(&newProduct).Error; err != nil {
		return ProductModel{}, err
	}

	return newProduct, nil
}

func (pq *ProductQuery) ListProduct() ([]ProductModel, error) {
	var products []ProductModel
	if err := pq.DB.Joins("JOIN user_models ON product_models.user_id = user_models.username").
		Select("product_models.*, user_models.nama as user_nama").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (pq *ProductQuery) GetProductByBarcode(barcode string) (ProductModel, error) {
	var product ProductModel
	if err := pq.DB.Where("barcode = ?", barcode).Joins("JOIN user_models ON product_models.user_id = user_models.username").
		Select("product_models.*, user_models.nama as user_nama").First(&product).Error; err != nil {
		return ProductModel{}, err
	}

	return product, nil
}

func (pq *ProductQuery) UpdateProduct(barcode string, updatedData ProductModel) (ProductModel, error) {
	var product ProductModel
	if err := pq.DB.Model(&product).Where("barcode = ?", barcode).Updates(updatedData).Error; err != nil {
		return ProductModel{}, err
	}

	return product, nil
}

func (pq *ProductQuery) DeleteProduct(barcode string) error {
	var product ProductModel
	if err := pq.DB.Model(&product).Where("barcode = ?", barcode).Delete(&product).Error; err != nil {
		return err
	}

	return nil
}
