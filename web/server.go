package web

import (
	"github.com/alekseinovikov/todo/repository"
	"github.com/alekseinovikov/todo/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routing(e)

	e.Logger.Fatal(e.Start(":8080"))
}

func routing(e *echo.Echo) {
	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepository)
	h := NewTodoWebHandler(todoService)

	e.POST("/api/todo", h.Create)
	e.GET("/api/todo/:id", h.GetById)
}
