package model

import (
	"time"

	"gorm.io/gorm"
)

type CustomerModel struct {
	Hp         string `gorm:"type:varchar(13);primaryKey"`
	Nama       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt   `gorm:"index"`
	Pembelians []PembelianModel `gorm:"foreignKey:CustomerID"`
}

type CustomerQuery struct {
	DB *gorm.DB
}

func (cq *CustomerQuery) AddCustomer(newCustomer CustomerModel) (CustomerModel, error) {
	if err := cq.DB.Create(&newCustomer).Error; err != nil {
		return CustomerModel{}, err
	}

	return newCustomer, nil
}

func (cq *CustomerQuery) ListCustomer() ([]CustomerModel, error) {
	var customers []CustomerModel
	if err := cq.DB.Find(&customers).Error; err != nil {
		return nil, err
	}

	return customers, nil
}

func (cq *CustomerQuery) GetCustomerByHp(hp string) (CustomerModel, error) {
	var customers CustomerModel
	if err := cq.DB.Where("hp = ?", hp).First(&customers).Error; err != nil {
		return CustomerModel{}, err
	}

	return customers, nil
}

func (cq *CustomerQuery) UpdateCustomer(hp string, updatedData CustomerModel) (CustomerModel, error) {
	var customers CustomerModel
	if err := cq.DB.Model(&customers).Where("hp = ?", hp).Updates(updatedData).Error; err != nil {
		return CustomerModel{}, err
	}

	return customers, nil
}

func (cq *CustomerQuery) DeleteCustomer(hp string) error {
	var customers CustomerModel
	if err := cq.DB.Model(&customers).Where("hp = ?", hp).Delete(&customers).Error; err != nil {
		return err
	}

	return nil
}
