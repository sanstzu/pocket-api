package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/sanstzu/pocket-api/src/api/handlers"
)

func UseSubmitRoutes(e *echo.Group) {
	e.POST("/submit", handlers.Submit)
}
