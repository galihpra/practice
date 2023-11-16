package model

import (
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	Username  string `gorm:"type:varchar(55);primaryKey"`
	Nama      string
	Status    int
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Products  []ProductModel `gorm:"foreignKey:UserID"`
}

type UserQuery struct {
	DB *gorm.DB
}

func (uq *UserQuery) Register(newUser UserModel) (UserModel, error) {
	if err := uq.DB.Create(&newUser).Error; err != nil {
		return UserModel{}, err
	}

	return newUser, nil
}

func (uq *UserQuery) Login(username string, password string) (UserModel, error) {
	var result = new(UserModel)

	if err := uq.DB.Where("username = ? AND password = ?", username, password).First(result).Error; err != nil {
		return UserModel{}, err
	}

	return *result, nil
}

func (uq *UserQuery) ListUser() ([]UserModel, error) {
	var users []UserModel
	if err := uq.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
