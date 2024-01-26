package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandleIndex(e echo.Context) error {
	return e.Render(http.StatusOK, "index", map[string]interface{}{"IsLoggedIn": isLoggedIn})
}
