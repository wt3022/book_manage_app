package routes

import (
	"book_manage_app/ent"
	"book_manage_app/internal/controller"
	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo, db *ent.Client) {
	g := e.Group("/users")

	g.GET("/", controller.UserGet)
}
