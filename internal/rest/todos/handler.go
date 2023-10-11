package todos

import (
	"com/khan/todo/internal/model"
	"com/khan/todo/internal/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

func getTodosHandler(c echo.Context) error {
	e.GET("/todos", func(c echo.Context) error {
		var todos []model.Todo
		DB.Find(&todos)

		if db.Error == nil {
			return c.JSON(http.StatusInternalServerError, response.Response{
				Success:  false,
				Code:     http.StatusInternalServerError,
				Messages: []string{"Failed to fetch todos"},
				Data:     nil,
			})
		}

		return c.JSON(http.StatusOK, response.Response{
			Success:  true,
			Code:     http.StatusOK,
			Messages: []string{},
			Data:     todos,
		})
	})
}
