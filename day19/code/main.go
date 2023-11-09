package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// // get user by id
func GetUserController(c echo.Context) error {
	// Extract user ID from request URL
	userID := c.Param("id")

	// Convert ID to integer
	id, err := strconv.Atoi(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid user ID",
		})
	}

	// Find the user with the specified ID
	var user *User
	for _, u := range users {
		if u.Id == id {
			user = &u // Assign the found user to the variable
			break
		}
	}

	if user == nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "User not found",
		})
	}

	// Return the found user as JSON
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get user",
		"user":    user,
	})
}

// // delete user by id
func DeleteUserController(c echo.Context) error {
	// Extract user ID from request URL
	userID := c.Param("id")

	// Convert ID to integer
	id, err := strconv.Atoi(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid user ID",
		})
	}

	// Find the user with the specified ID
	userIndex := -1
	for i, u := range users {
		if u.Id == id {
			userIndex = i // Assign the index of the found user to the variable
			break
		}
	}

	if userIndex == -1 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "User not found",
		})
	}

	// Delete the user from the array
	users = append(users[:userIndex], users[userIndex+1:]...)

	// Return a success message
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success delete user",
	})
}

// // update user by id
func UpdateUserController(c echo.Context) error {
	// Extract user ID from request URL
	userID := c.Param("id")

	// Convert ID to integer
	id, err := strconv.Atoi(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid user ID",
		})
	}

	// Find the user with the specified ID
	var user *User
	for _, u := range users {
		if u.Id == id {
			user = &u // Assign the found user to the variable
			break
		}
	}

	if user == nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "User not found",
		})
	}

	// Bind the update data from the request
	err = c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid update data",
		})
	}

	// Update the user
	users[id-1] = *user

	// Return a success message
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success update user",
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.GET("/users/:id", GetUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
