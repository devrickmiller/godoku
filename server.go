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
		m.MockMat(solve.HARDEST, 1)
		err := m.Solve()
		if err != nil {
			return err
		}		
		return c.String(200, m.String() + fmt.Sprintf("\nToggle Count = %d\n", solve.ToggleCount)) 
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