package main

import (
	"day-22/auth"
	"day-22/config"
	"day-22/customer"
	"day-22/model"
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

	var auth = auth.AuthSystem{DB: db}
	var cust = customer.CustomerSystem{DB: db}

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

	e.Logger.Fatal(e.Start(":8000"))
}
