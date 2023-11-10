package web

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetAllBooks(t *testing.T) {
	e := echo.New()

	t.Run("no todo by id - should return 404", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/todo/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		h := &handler{}
		_ = h.GetById(c)

		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Empty(t, rec.Body.String())
	})

	t.Run("post todo - should return 200 and response with id", func(t *testing.T) {
		todoJson := buildTodoJson("")
		req := httptest.NewRequest(http.MethodPost, "/api/todo", strings.NewReader(todoJson))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		h := &handler{}
		_ = h.Create(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, todoJson, buildTodoJson("1"))
	})
}

func buildTodoJson(id string) string {
	if id == "" {
		return `{"title":"test","description":"test","status":"test"}`
	}

	return `{"id":"` + id + `","title":"test","description":"test","status":"test"}`
}
