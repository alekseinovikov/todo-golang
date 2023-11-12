package web

import (
	"github.com/alekseinovikov/todo/tests/todotest"
	"github.com/labstack/echo/v4"
	"net/http"
	"testing"
)

func TestGetAllBooks(t *testing.T) {
	e := echo.New()

	t.Run("no todo by id - should return 404", func(t *testing.T) {
		todotest.NewTodoRequest(t, e, &handler{}).
			SetId("1").
			Run().
			AssertStatus(http.StatusNotFound).
			AssertBody("")
	})

	t.Run("post todo - should return 200 and response with id", func(t *testing.T) {
		todotest.NewTodoCreateRequest(t, e, &handler{}).
			SetRequest(todotest.TodoCreateRequestBody{
				Title:       "test",
				Description: "test",
				Status:      "test",
			}).
			Run().
			AssertStatus(http.StatusOK).
			AssertBody(todotest.TodoCreateResponseBody{
				Id:          "1",
				Title:       "test",
				Description: "test",
				Status:      "test",
			})
	})
}
