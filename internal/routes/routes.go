package routes

import (
	"github.com/Tanachuns/teebaan-payment/internal/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/users/:id", controllers.GetUser)
}
