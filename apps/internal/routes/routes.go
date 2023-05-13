package routes

import (
	"book_manage_app/ent"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, db *ent.Client) {
	UserRoutes(e, db)
}
