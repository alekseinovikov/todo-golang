package web

import (
	domain "github.com/alekseinovikov/todo/domain"
	"github.com/labstack/echo/v4"
	"net/http"
)

type handler struct {
}

func (h *handler) GetById(c echo.Context) error {
	return c.String(http.StatusNotFound, "")
}

func (h *handler) Create(c echo.Context) error {
	var todo domain.Todo
	err := c.Bind(&todo)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	todo.Id = "1"
	return c.JSON(http.StatusOK, todo)
}
