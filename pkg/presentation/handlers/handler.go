package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	return c.Render(http.StatusOK, "top.html", map[string]interface{}{
		"title": "Top Page",
		"h1": 	"タスク一覧",
	})
}

func About(c echo.Context) error {
	return c.Render(http.StatusOK, "about.html", map[string]interface{}{
		"title": "About",
		"h1": 	"About Page",
	})
}
