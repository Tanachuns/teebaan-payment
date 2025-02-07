package routes

import (
	"github.com/Tanachuns/teebaan-payment/internal/controllers"
	"github.com/Tanachuns/teebaan-payment/internal/services"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api_group := e.Group("api/v1/")

	api_group.GET("users/:id", controllers.GetUser, services.TokenMiddleware)
	e.POST("token", controllers.GetToken)
}
