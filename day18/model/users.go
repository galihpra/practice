package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	HP       string
	Nama     string
	Alamat   string
	Password string
	Todos    []Todo
}
