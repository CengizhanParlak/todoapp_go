package todos

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", GetIndexHandler)
	e.GET("/todos", GetTodosHandler)
	e.POST("/todos", CreateTodoHandler)
	e.PUT("/todos/:id", UpdateTodoStatusHandler)
	e.DELETE("/todos/:id", DeleteTodoHandler)
}
