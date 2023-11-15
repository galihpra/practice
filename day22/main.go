package main

import (
	"day-22/auth"
	"day-22/config"
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

	var auth = auth.AuthSystem{DB: db}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/login", auth.LoginHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}
