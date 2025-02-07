package main

import (
	"github.com/Tanachuns/teebaan-payment/internal/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func main() {
	e := echo.New()
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(1))))
	// Register routes
	routes.RegisterRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":1234"))
}
