package handlers

import (
	"da1ch1a/go-todo/pkg/presentation/registry"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// TODO ハンドラが増えたら抽象化してそこにregistryを渡すようにする
type TaskHandler struct {
	Registry *registry.Registry
}

func (h *TaskHandler) List(c echo.Context) error {
	accept := c.Request().Header.Get(echo.HeaderAccept)
	if accept != echo.MIMEApplicationJSON {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid Accept header: %s", accept))
	}

	return c.JSON(http.StatusOK, h.Registry.TaskUsecase.List())
}

func (h *TaskHandler) Create(c echo.Context) error {
	accept := c.Request().Header.Get(echo.HeaderAccept)
	if accept != echo.MIMEApplicationJSON {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid Accept header: %s", accept))
	}

	err := h.Registry.TaskUsecase.Create(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
