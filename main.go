package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func main() {
	e := echo.New()
	e.Debug = true

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
		rate.Limit(20),
	)))

	NewTemplateRenderer(e, "public/*.*html")

	e.GET("/", HandleIndex)
	e.GET("/data", HandleData)
	e.GET("/data/:id", HandleDataById)
	e.PUT("/data/:id", HandleDataUpdateById)
	e.GET("/data/:id/edit", HandleDataEdit)
	e.POST("/login", HandleLogin)
	e.POST("/logout", HandleLogout)

	e.Logger.Fatal(e.Start("localhost:4040"))
}
