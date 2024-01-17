package api

import (
	"github.com/labstack/echo/v4"
	"github.com/sanstzu/pocket-api/src/api/routes"
)

func Start() {
	e := echo.New()

	routes.UseSubmitRoutes(e.Group("/api"))

	e.Logger.Fatal(e.Start(":8080"))
}
