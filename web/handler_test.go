package web

import (
	"github.com/alekseinovikov/todo/repository"
	"github.com/alekseinovikov/todo/service"
	"github.com/alekseinovikov/todo/tests/todotest"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"testing"
)

func TestGetAllBooks(t *testing.T) {
	e := echo.New()
	e.Use(middleware.Logger())

	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepository)
	handler := NewTodoWebHandler(todoService)

	t.Run("no todo by id - should return 404", func(t *testing.T) {
		todotest.NewTodoRequest(t, e, handler).
			SetId("1").
			Run().
			AssertStatus(http.StatusNotFound).
			AssertPayload("todo not found")
	})

	t.Run("post todo - should return 200 and response with id", func(t *testing.T) {
		todotest.NewTodoCreateRequest(t, e, handler).
			SetRequest(todotest.TodoCreateRequestBody{
				Title:       "test",
				Description: "test",
				Status:      "test",
			}).
			Run().
			AssertStatus(http.StatusOK).
			AssertIdIsNotEmpty().
			AssertBody(todotest.TodoCreateRequestBody{
				Title:       "test",
				Description: "test",
				Status:      "test",
			})
	})

	t.Run("post todo - get by created id - returns the whole todo", func(t *testing.T) {
		createdId := todotest.NewTodoCreateRequest(t, e, handler).
			SetRequest(todotest.TodoCreateRequestBody{
				Title:       "test",
				Description: "test",
				Status:      "test",
			}).
			Run().
			AssertStatus(http.StatusOK).
			AssertIdIsNotEmpty().
			AssertBody(todotest.TodoCreateRequestBody{
				Title:       "test",
				Description: "test",
				Status:      "test",
			}).
			GetId()

		todotest.NewTodoRequest(t, e, handler).
			SetId(createdId).
			Run().
			AssertStatus(http.StatusOK).
			AssertBody(todotest.TodoResponseBody{
				Id:          createdId,
				Title:       "test",
				Description: "test",
				Status:      "test",
			})
	})
}
