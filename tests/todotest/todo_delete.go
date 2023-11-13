package todotest

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type todoDelete interface {
	Delete(c echo.Context) error
}

type TodoDeleteRequest struct {
	id       string
	echo     *echo.Echo
	recorder *httptest.ResponseRecorder
	deleter  todoDelete
	t        *testing.T
}

func NewTodoDelete(t *testing.T, e *echo.Echo, deleter todoDelete) *TodoDeleteRequest {
	rec := httptest.NewRecorder()
	return &TodoDeleteRequest{t: t, recorder: rec, echo: e, deleter: deleter}
}

func (r *TodoDeleteRequest) SetId(id string) *TodoDeleteRequest {
	r.id = id
	return r
}

func (r *TodoDeleteRequest) Run() *TodoDeleteRequest {
	req := httptest.NewRequest(http.MethodDelete, "/api/todo", nil)
	c := r.echo.NewContext(req, r.recorder)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues(r.id)

	_ = r.deleter.Delete(c)
	return r
}

func (r *TodoDeleteRequest) AssertStatus(status int) *TodoDeleteRequest {
	assert.Equal(r.t, status, r.recorder.Code)
	return r
}

func (r *TodoDeleteRequest) AssertPayload(body string) *TodoDeleteRequest {
	assert.Equal(r.t, body, r.recorder.Body.String())
	return r
}
