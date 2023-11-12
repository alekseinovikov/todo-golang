package todotest

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type todoById interface {
	GetById(c echo.Context) error
}

type TodoRequest struct {
	id       string
	echo     *echo.Echo
	recorder *httptest.ResponseRecorder
	getter   todoById
	t        *testing.T
}

func NewTodoRequest(t *testing.T, e *echo.Echo, getter todoById) *TodoRequest {
	rec := httptest.NewRecorder()
	return &TodoRequest{t: t, recorder: rec, echo: e, getter: getter}
}

func (r *TodoRequest) SetId(id string) *TodoRequest {
	r.id = id
	return r
}

func (r *TodoRequest) Run() *TodoRequest {
	req := httptest.NewRequest(http.MethodGet, "/api/todo/"+r.id, nil)
	c := r.echo.NewContext(req, r.recorder)
	_ = r.getter.GetById(c)
	return r
}

func (r *TodoRequest) AssertStatus(status int) *TodoRequest {
	assert.Equal(r.t, status, r.recorder.Code)
	return r
}

func (r *TodoRequest) AssertBody(body string) *TodoRequest {
	assert.Equal(r.t, body, r.recorder.Body.String())
	return r
}
