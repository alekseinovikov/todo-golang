package web

import "github.com/labstack/echo/v4"

func Run() {
	e := echo.New()

	routing(e)

	e.Logger.Fatal(e.Start(":8080"))
}

func routing(e *echo.Echo) {
	h := &handler{}
	g := e.Group("/api/todo")

	g.GET("/:id", h.GetById)
}
