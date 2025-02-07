package controllers

import (
	"net/http"
	"strconv"

	"github.com/Tanachuns/teebaan-payment/internal/models"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user, err := models.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}
	return c.JSON(http.StatusOK, user)
}
