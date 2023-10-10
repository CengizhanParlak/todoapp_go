package main

import (
	"com/khan/todo/internal/model"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func main() {
	db := initializeDb()
	defer db.DB()
	db.AutoMigrate(&model.Todo{})

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.POST("/todos", func(c echo.Context) error {
		todo := new(model.Todo)

		if err := c.Bind(todo); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		todo.TodoStatus = model.Incomplete
		todo.CreatedAt = time.Now()
		todo.UpdatedAt = time.Now()

		db.Create(&todo)

		if db.Error != nil {
			return c.JSON(http.StatusInternalServerError, db.Error)
		}

		return c.JSON(http.StatusOK, todo)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func initializeDb() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=postgres dbname=todoapp_db port=5432 sslmode=disable TimeZone=Europe/Istanbul",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	return db
}
