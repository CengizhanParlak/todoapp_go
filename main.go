package main

import (
	"com/khan/todo/internal/db"
	"com/khan/todo/internal/model"
	"com/khan/todo/internal/rest/todos"
	"github.com/labstack/echo/v4"
)

func main() {
	err := db.InitDb()
	if err != nil {
		panic("Failed to connect to the database")
	}

	err = db.DB.AutoMigrate(&model.Todo{})

	if err != nil {
		panic("Failed to migrate the database")
	}

	e := echo.New()
	todos.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
