package attendance

import (
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/", home)
	e.GET("/attendance", list)
	e.POST("/attendance", create)
}
