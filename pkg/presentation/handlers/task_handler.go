package handlers

import (
	"da1ch1a/go-todo/pkg/application/usecase"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {}

func (h *TaskHandler) List(c echo.Context) error {
	accept := c.Request().Header.Get(echo.HeaderAccept)
	if accept == echo.MIMEApplicationJSON {
		taskUsecase := usecase.TaskUsecase{}
		return taskUsecase.List(c)
	}

	return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid Accept header: %s", accept))
}

func (h *TaskHandler) Json(e echo.Context) error {
	header := e.Request().Header
	accept := header.Get(echo.HeaderAccept)

	if accept == echo.MIMEApplicationJSON {
		return e.JSON(http.StatusOK, map[string]string{"response": "JSONで返すぜ"})
	}

	err := echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid Accept header: %s", accept))

	return err
}
