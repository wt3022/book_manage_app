package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Tags users
// @Router /users/ [get]
func UserGet(c echo.Context) error {
	return c.String(http.StatusOK, "hoge")
}
