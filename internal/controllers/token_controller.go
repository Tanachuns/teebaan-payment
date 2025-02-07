package controllers

import (
	"net/http"

	"github.com/Tanachuns/teebaan-payment/internal/services"
	"github.com/labstack/echo/v4"
)

type RequestBody struct {
	APIKey string `json:"api_key"`
}

func GetToken(c echo.Context) error {
	var req RequestBody
	if err := c.Bind(&req); err != nil || req.APIKey == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid or empty request"})
	}

	// Access API Key
	apiKey := req.APIKey

	token, err := services.GenerateToken(apiKey)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})

}
