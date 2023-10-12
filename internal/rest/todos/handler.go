package todos

import (
	"com/khan/todo/internal/db"
	"com/khan/todo/internal/model"
	"com/khan/todo/internal/response"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func GetIndexHandler(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Welcome to the todos API")
}

func GetTodosHandler(ctx echo.Context) error {
	var todos []model.Todo
	db.DB.Find(&todos)
	return ctx.JSON(http.StatusOK, response.Response{
		Success:  true,
		Code:     http.StatusOK,
		Data:     todos,
		Messages: []string{},
	})
}

func UpdateTodoStatusHandler(ctx echo.Context) error {
	id := ctx.Param("id")
	var todo model.Todo
	db.DB.First(&todo, id)

	if todo.ID == 0 {
		return ctx.JSON(http.StatusNotFound, response.Response{
			Success:  false,
			Code:     http.StatusNotFound,
			Data:     nil,
			Messages: []string{"Todo not found"},
		})
	}

	if todo.TodoStatus == model.Completed {
		todo.TodoStatus = model.Incomplete
	} else {
		todo.TodoStatus = model.Completed
	}
	todo.UpdatedAt = time.Now()
	db.DB.Save(&todo)

	if db.DB.Error != nil {
		return ctx.JSON(http.StatusInternalServerError, response.Response{
			Success:  false,
			Code:     http.StatusInternalServerError,
			Data:     nil,
			Messages: []string{db.DB.Error.Error()},
		})
	}

	return ctx.JSON(http.StatusOK, response.Response{
		Success:  true,
		Code:     http.StatusOK,
		Data:     todo,
		Messages: []string{},
	})
}

func CreateTodoHandler(ctx echo.Context) error {
	todo := new(model.Todo)

	if err := ctx.Bind(todo); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	todo.TodoStatus = model.Incomplete
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()

	db.DB.Create(&todo)

	if db.DB.Error != nil {
		return db.DB.Error
	}

	return ctx.JSON(http.StatusOK, response.Response{
		Success:  true,
		Code:     http.StatusOK,
		Data:     todo,
		Messages: []string{},
	})
}

func DeleteTodoHandler(ctx echo.Context) error {
	id := ctx.Param("id")
	var todo model.Todo
	db.DB.First(&todo, id)

	if todo.ID == 0 {
		return ctx.JSON(http.StatusNotFound, response.Response{
			Success:  false,
			Code:     http.StatusNotFound,
			Data:     nil,
			Messages: []string{"Todo not found"},
		})
	}

	db.DB.Delete(&todo)

	if db.DB.Error != nil {
		return ctx.JSON(http.StatusInternalServerError, response.Response{
			Success:  false,
			Code:     http.StatusInternalServerError,
			Data:     nil,
			Messages: []string{db.DB.Error.Error()},
		})
	}

	return ctx.JSON(http.StatusOK, response.Response{
		Success:  true,
		Code:     http.StatusOK,
		Data:     nil,
		Messages: []string{},
	})
}
