package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// https://echo.labstack.com/docs/error-handling
func ErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := "Internal Server Error"
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message.(string)
	}
	c.Logger().Error(err)
	if err := c.JSON(code, map[string]interface{}{
		"message": message,
	}); err != nil {
		c.Logger().Error(err)
	}
}
