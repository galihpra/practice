package todo

import (
	"fmt"
	"sqlgo/model"
	"time"

	"gorm.io/gorm"
)

type TodoSystem struct {
	DB *gorm.DB
}

func (ts *TodoSystem) Add(userID uint) (model.Todo, bool) {
	var newTodo = new(model.Todo)
	var inputTanggal string
	fmt.Print("Masukkan Nama Kegiatan")
	fmt.Scanln(&newTodo.Nama)
	fmt.Print("Masukkan Deskripsi")
	fmt.Scanln(&newTodo.Deskripsi)
	fmt.Print("Masukkan Deadline")
	fmt.Scanln(&inputTanggal)
	var dateFormat = "2006-01-02"
	newTodo.Deadline, _ = time.Parse(dateFormat, inputTanggal)
	newTodo.UserID = userID

	err := ts.DB.Create(newTodo).Error
	if err != nil {
		fmt.Println("input error:", err.Error())
		return model.Todo{}, false
	}

	return *newTodo, true
}

func (ts *TodoSystem) Read(userID uint) ([]model.Todo, bool) {
	var todos []model.Todo

	// Query todos for the given user ID.
	qry := ts.DB.Where("user_id = ?", userID).Find(&todos)
	err := qry.Error

	if err != nil {
		fmt.Println("Error read data table:", err.Error())
		return nil, false
	}

	return todos, true
}

func (ts *TodoSystem) Delete(userID uint, namaTodo string) bool {
	var todo model.Todo
	qry := ts.DB.Where("user_id = ? AND nama = ?", userID, namaTodo).First(&todo)
	if qry.Error != nil {
		fmt.Println("Kegiatan tidak ditemukan")
		return false
	}

	if err := ts.DB.Delete(&todo).Error; err != nil {
		fmt.Println("Gagal menghapus todo:", err.Error())
		return false
	}

	return true
}

func (ts *TodoSystem) Update(userID uint, namaTodo, newNamaTodo string) bool {
	var todo model.Todo
	qry := ts.DB.Where("user_id = ? AND nama = ?", userID, namaTodo).First(&todo)
	if qry.Error != nil {
		fmt.Println("Kegiatan tidak ditemukan")
		return false
	}

	todo.Nama = newNamaTodo
	if err := ts.DB.Save(&todo).Error; err != nil {
		fmt.Println("Gagal mengupdate todo:", err.Error())
		return false
	}

	return true
}
