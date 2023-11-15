package main

import (
	"day-22/auth"
	"day-22/config"
	"day-22/customer"
	detailpembelian "day-22/detail_pembelian"
	"day-22/model"
	"day-22/pembelian"
	"day-22/products"
	"day-22/users"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		fmt.Println("Something happened", err.Error())
		return
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Customer{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Product{})
	db.AutoMigrate(&model.Pembelian{})
	db.AutoMigrate(&model.DetailPembelian{})

	var auth = auth.AuthSystem{DB: db}
	var cust = customer.CustomerSystem{DB: db}
	var user = users.UserSystem{DB: db}
	var product = products.ProductSystem{DB: db}
	var pembelian = pembelian.PembelianSystem{DB: db}
	var detail = detailpembelian.DetailPembelianSystem{DB: db}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Login
	e.POST("/login", auth.LoginHandler)

	// Customer
	e.POST("/customers", cust.CreateCustomer)
	e.GET("/customers", cust.ReadCustomer)
	e.DELETE("/customers/:hp", cust.DeleteCustomer)
	e.PUT("/customers/:hp", cust.UpdateCustomer)

	// User
	e.POST("/users", user.CreateUser)
	e.GET("/users", user.ReadUser)
	e.GET("/users/:username", user.ReadUserById)
	e.DELETE("/users/:username", user.DeleteUser)
	e.PUT("/users/:username", user.UpdateUser)

	// Product
	e.POST("/products", product.CreateProduct)
	e.GET("/products", product.ReadProduct)
	e.GET("/products/:barcode", product.ReadProductById)
	e.DELETE("/products/:barcode", product.DeleteProduct)
	e.PUT("/products/:barcode", product.UpdateProduct)

	// Pembelian
	e.POST("/transactions", pembelian.CreatePembelian)
	e.GET("/transactions", pembelian.ReadPembelian)
	e.GET("/transactions/:no_invoice", pembelian.ReadPembelianById)
	e.DELETE("/transactions/:no_invoice", pembelian.DeletePembelian)
	e.PUT("/transactions/:no_invoice", pembelian.UpdatePembelian)

	// Detail Pembelian
	e.POST("/detail_transactions", detail.CreateDetailPembelian)
	e.GET("/detail_transactions", detail.ReadDetailPembelian)
	e.GET("/detail_transactions/:pembelian_id", detail.ReadDetailPembelianById)
	e.DELETE("/detail_transactions/:pembelian_id", detail.DeleteDetailPembelian)
	e.PUT("/detail_transactions/:pembelian_id", detail.UpdateDetailPembelian)

	e.Logger.Fatal(e.Start(":8000"))
}
