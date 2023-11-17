package main

import (
	"day-22/config"
	"day-22/controller/customer"
	detailpembelian "day-22/controller/detail_pembelian"
	"day-22/controller/pembelian"
	"day-22/controller/product"
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

	db.AutoMigrate(model.UserModel{}, &model.ProductModel{}, &model.CustomerModel{}, &model.PembelianModel{}, &model.DetailPembelianModel{})

	modelUser := model.UserQuery{DB: db}
	userController := user.UserController{Model: modelUser}

	modelCustomer := model.CustomerQuery{DB: db}
	customerController := customer.CustomerController{Model: modelCustomer}

	modelProduct := model.ProductQuery{DB: db}
	productController := product.ProductController{Model: modelProduct}

	modelPembelian := model.PembelianQuery{DB: db}
	pembelianController := pembelian.PembelianController{Model: modelPembelian}

	modelDetailPembelian := model.DetailPembelianQuery{DB: db}
	detailPembelianController := detailpembelian.DetailPembelianController{Model: modelDetailPembelian}

	routes.InitRoute(e, userController, customerController, productController, pembelianController, detailPembelianController)

	e.Logger.Fatal(e.Start(":8000"))
}
