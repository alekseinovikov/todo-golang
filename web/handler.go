package web

import (
	domain "github.com/alekseinovikov/todo/domain"
	"github.com/labstack/echo/v4"
	"net/http"
)

type todoService interface {
	GetById(id string) (error, domain.Todo)
	Create(todo domain.Todo) (error, domain.Todo)
	Delete(id string) error
}

type TodoHandler struct {
	service todoService
}

func NewTodoWebHandler(service todoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) GetById(c echo.Context) error {
	id := c.Param("id")
	err, todo := h.service.GetById(id)
	if err == nil {
		return c.JSON(http.StatusOK, todo)
	}

	if err == domain.ErrTodoNotFound {
		return c.String(http.StatusNotFound, err.Error())
	}

	return c.String(http.StatusInternalServerError, err.Error())
}

func (h *TodoHandler) Create(c echo.Context) error {
	var todo domain.Todo
	err := c.Bind(&todo)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err, todo = h.service.Create(todo)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	err := h.service.Delete(id)
	if err != nil {
		if err == domain.ErrTodoNotFound {
			return c.String(http.StatusNotFound, err.Error())
		}

		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "OK")
}
