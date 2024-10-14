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

func Json(e echo.Context) error {
	header := e.Request().Header
	accept := header.Get(echo.HeaderAccept)

	if accept == echo.MIMEApplicationJSON {
		return e.JSON(http.StatusOK, map[string]string{"response": "JSONで返すぜ"})
	}

	err := echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid Accept header: %s", accept))

	return err
}

