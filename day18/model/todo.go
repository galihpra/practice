package model

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Nama      string
	Deskripsi string
	Deadline  time.Time
	IsDone    bool
	UserID    uint
}
