package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// struct user
type user struct {
	ID   int    `json : "id"`
	name string `json : "name"`
}

// initial data
var users = map[int]*user{}
var count int = 0

// ======
// Handlers
// ======
func createUser(c echo.Context) error {
	// membuat stuct data user
	data := &user{
		ID: count,
	}
	//  mengambil data dari req
	if err := c.Bind(data); err != nil {
		return err
	}
	// menambahkan user
	users[data.ID] = data
	count++

	// response
	return c.JSON(http.StatusCreated, data)
}

func getAllUsers(c echo.Context) error {
	// response
	return c.JSON(http.StatusOK, users)
}

func getUserByID(c echo.Context) error {
	// mengambil params dan convert ke int
	id, _ := strconv.Atoi(c.Param("id"))
	// mengecek user ada atau tidak
	if _, isExist := users[id]; isExist == false {
		return c.String(http.StatusNotFound, "User Not Found")
	}
	// response
	return c.JSON(http.StatusOK, users[id])
}

func updateUser(c echo.Context) error {
	// membuat struct user
	data := new(user)
	// mengambil req body
	if err := c.Bind(data); err != nil {
		return err
	}
	// mengambil params dan convert ke int
	id, _ := strconv.Atoi(c.Param("id"))
	// mengecek user ada atau tidak
	if _, isExist := users[id]; isExist == false {
		return c.String(http.StatusNotFound, "User Not Found")
	}
	// mengubah data
	users[id].name = data.name
	// response
	return c.JSON(http.StatusOK, users[id])
}

func deleteUser(c echo.Context) error {
	// mengambil params dan convert ke int
	id, _ := strconv.Atoi(c.Param("id"))
	// mengecek user ada atau tidak
	if _, isExist := users[id]; isExist == false {
		return c.String(http.StatusNotFound, "User Not Found")
	}
	// menghapus user param 1 = map , 2 = key value
	delete(users, id)
	// response
	return c.NoContent(http.StatusNoContent)

}

func hello(c echo.Context) error {
	// response
	return c.String(http.StatusOK, "Hello Mas Bro")
}

// Main Fuction
func main() {
	e := echo.New()

	// middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// routes
	e.GET("/", hello)
	e.GET("/users", getAllUsers)
	e.GET("/user/:id", getUserByID)
	e.POST("/user", createUser)
	e.PUT("/user/:id", updateUser)
	e.DELETE("/user/:id", deleteUser)

	// start server
	e.Logger.Fatal(e.Start(":5000"))
}
