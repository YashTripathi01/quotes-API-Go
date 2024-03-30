package main

import (
	"quotes-api/handlers"
	"quotes-api/initializers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handlers.Home)
	e.GET("/quotes", handlers.GetQuotesList)
	e.GET("/quotes/today", handlers.GetQuoteOfTheDayHandler)
	e.POST("/quotes", handlers.CreateQuotes)
	e.PUT("/quotes", handlers.ResetQuoteOfTheDay)

	e.Logger.Fatal(e.Start(":1323"))
}
