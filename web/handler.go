package web

import (
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

type handler struct {
}

func (h *handler) GetById(c echo.Context) error {
	return c.String(http.StatusNotFound, "")
}

func (h *handler) Create(c echo.Context) interface{} {
	all, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSONBlob(http.StatusOK, all)
}
