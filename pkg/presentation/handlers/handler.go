package handlers

import (
	"da1ch1a/go-todo/pkg/application/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	taskUsecase := usecase.TaskUsecase{}
	tasks := taskUsecase.List()

	return c.Render(http.StatusOK, "top.html", map[string]interface{}{
		"H1":    "タスク一覧",
		"Tasks": &tasks,
	})
}

func About(c echo.Context) error {
	return c.Render(http.StatusOK, "about.html", map[string]interface{}{
		"h1": "About Page",
	})
}
