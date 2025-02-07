package routes

import (
	"github.com/Tanachuns/teebaan-payment/internal/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api_group := e.Group("api/v1/")
	api_group.GET("users/:id", controllers.GetUser)
}
