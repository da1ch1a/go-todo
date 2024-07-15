package routers

import (
	"da1ch1a/go-todo/pkg/presentation/handlers"

	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) {
	e.GET("/", handlers.Home)
	e.GET("/about", handlers.About)
}
