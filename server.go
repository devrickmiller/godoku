package main

import (
	"fmt"
	"os"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello from Godoku! <3")	
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})	
	})

	httpPort := os.Getenv("PORT")
	fmt.Println("PORT = ", httpPort)

	if httpPort == "" { httpPort = "8080" }

	e.Logger.Fatal(e.Start(":" + httpPort))
	//e.Logger.Fatal(e.Start(":8080"))
}