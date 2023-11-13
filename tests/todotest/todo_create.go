package todotest

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

type todoCreate interface {
	Create(c echo.Context) error
}
type TodoCreateRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type TodoCreateResponseBody struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type TodoCreateRequest struct {
	t        *testing.T
	echo     *echo.Echo
	recorder *httptest.ResponseRecorder
	creator  todoCreate
	request  TodoCreateRequestBody
}

func NewTodoCreateRequest(t *testing.T, e *echo.Echo, creator todoCreate) *TodoCreateRequest {
	rec := httptest.NewRecorder()
	return &TodoCreateRequest{t: t, recorder: rec, echo: e, creator: creator}
}

func (r *TodoCreateRequest) SetRequest(request TodoCreateRequestBody) *TodoCreateRequest {
	r.request = request
	return r
}

func (r *TodoCreateRequest) Run() *TodoCreateRequest {
	payload, _ := json.Marshal(r.request)
	req := httptest.NewRequest("POST", "/api/todo", bytes.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c := r.echo.NewContext(req, r.recorder)
	_ = r.creator.Create(c)
	return r
}

func (r *TodoCreateRequest) AssertStatus(status int) *TodoCreateRequest {
	assert.Equal(r.t, status, r.recorder.Code)
	return r
}

func (r *TodoCreateRequest) AssertBody(body TodoCreateRequestBody) *TodoCreateRequest {
	var actualBody TodoCreateRequestBody
	_ = json.Unmarshal(r.recorder.Body.Bytes(), &actualBody)
	assert.Equal(r.t, body, actualBody)
	return r
}

func (r *TodoCreateRequest) AssertIdIsNotEmpty() *TodoCreateRequest {
	var actualBody TodoCreateResponseBody
	_ = json.Unmarshal(r.recorder.Body.Bytes(), &actualBody)
	assert.NotEmpty(r.t, actualBody.Id)
	return r
}

func (r *TodoCreateRequest) GetId() string {
	var actualBody TodoCreateResponseBody
	_ = json.Unmarshal(r.recorder.Body.Bytes(), &actualBody)
	return actualBody.Id
}
