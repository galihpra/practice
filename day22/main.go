package main

import (
	"day-22/config"
	"day-22/controller/customer"
	"day-22/controller/user"
	"day-22/model"
	"day-22/routes"
	"day-22/utils/database"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	cfg := config.InitConfig()

	if cfg == nil {
		e.Logger.Fatal("tidak bisa start karena ENV error")
		return
	}

	db, err := database.InitMySQL(*cfg)

	if err != nil {
		e.Logger.Fatal("tidak bisa start karena DB error", err.Error())
		return
	}

	db.AutoMigrate(model.UserModel{}, &model.ProductModel{}, &model.CustomerModel{}, &model.PembelianModel{})

	modelUser := model.UserQuery{DB: db}
	userController := user.UserController{Model: modelUser}

	modelCustomer := model.CustomerQuery{DB: db}
	customerController := customer.CustomerController{Model: modelCustomer}

	routes.InitRoute(e, userController, customerController)

	e.Logger.Fatal(e.Start(":8000"))
}
