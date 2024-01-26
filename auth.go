package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

var isLoggedIn = false

func handleAuthentication(c echo.Context, loggedInState bool) error {
	isLoggedIn = loggedInState
	c.Response().Header().Set("HX-Redirect", "/")
	c.Response().WriteHeader(http.StatusFound)
	return c.JSON(http.StatusFound, "Redirect")
}

func HandleLogin(c echo.Context) error {
	return handleAuthentication(c, true)
}

func HandleLogout(c echo.Context) error {
	return handleAuthentication(c, false)
}
