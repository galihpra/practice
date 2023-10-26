package main

import (
	"fmt"
	"sqlgo/auth"
	"sqlgo/config"
	"sqlgo/model"
	"sqlgo/todo"
)

func main() {
	var inputMenu int
	db, err := config.InitDB()
	if err != nil {
		fmt.Println("Something happened", err.Error())
		return
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Todo{})

	var auth = auth.AuthSystem{DB: db}
	var todo = todo.TodoSystem{DB: db}

	for {
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("99. Exit")
		fmt.Print("Masukkan pilihan:")
		fmt.Scanln(&inputMenu)
		switch inputMenu {
		case 1:
			var menuLogin int
			result, permit := auth.Login()
			if permit {
				fmt.Println("Selamat datang ", result.Nama)
				for permit {
					fmt.Println("1. Add Todo")
					fmt.Println("2. Show List Todo")
					fmt.Println("3. Update Todo")
					fmt.Println("4. Hapus Todo")
					fmt.Println("0. Logout")
					fmt.Println("99. Exit")
					fmt.Print("Masukkan pilihan:")
					fmt.Scanln(&menuLogin)
					switch menuLogin {
					case 1:
						result, permit := todo.Add(result.ID)
						if permit {
							fmt.Println(result)
						}
					case 2:
						result, permit := todo.Read(result.ID)
						if permit {
							for _, a := range result {
								fmt.Println(a.Nama)
							}
						}
					case 3:
						var namaTodo string
						var newNamaTodo string
						fmt.Print("Masukkan nama kegiatan yang akan di update: ")
						fmt.Scanln(&namaTodo)
						fmt.Print("Masukkan kegiatan yang baru: ")
						fmt.Scanln(&newNamaTodo)
						success := todo.Update(result.ID, namaTodo, newNamaTodo)
						if success {
							fmt.Println("Kegiatan berhasil di update")
						}
					case 4:
						var namaTodo string
						fmt.Print("Masukkan nama kegiatan: ")
						fmt.Scanln(&namaTodo)
						success := todo.Delete(result.ID, namaTodo)
						if success {
							fmt.Println("Kegiatan berhasil dihapus")
						}
					case 0:
						permit = false
					case 99:
						fmt.Println("Thank you....")
						return

					}
				}
			}
		case 2:
			result, permit := auth.Register()
			if permit {
				fmt.Println(result)
			}
		case 99:
			fmt.Println("Thank you....")
			return
		default:
		}
	}

}
