package main

import (
	"webapp/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.GET("/", controllers.HomeHandler)
	e.GET("/home", controllers.HomeHandler)

	staticGroup := e.Group("/static")
	staticGroup.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "static",
		Browse: true,
	}))

	e.Logger.Fatal(e.Start(":4444"))
}
