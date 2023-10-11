package main

import (
	"fmt"
	"os"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"godoku/solve"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		m := solve.NewMat()
		m.MockMat()
		ints := m.GetPossibleSquareValues(8,1)
		fmt.Println(ints)
		//return c.HTML(http.StatusOK, "Hello from Godoku! <3")
		return c.String(200, m.String())
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