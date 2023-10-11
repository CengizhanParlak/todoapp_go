package main

import (
	"com/khan/todo/internal/model"
	"database/sql"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"time"
)

var DB *sql.DB

func main() {
	DB := initializeDb()
	defer DB.DB()
	err := DB.AutoMigrate(&model.Todo{})

	if err != nil {
		panic("Failed to migrate the database")
	}

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

		DB.Create(&todo)

		if DB.Error != nil {
			return c.JSON(http.StatusInternalServerError, DB.Error)
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
