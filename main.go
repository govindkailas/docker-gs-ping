package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/thanhpk/randstr"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	greetings := os.Getenv("GREETINGS")
	if greetings == "" {
		greetings = randstr.String(16)
	}

	//function to set the greetings to random if its null
	// func get_greetings() {
	// 	if greetings == "" {
	// 		greetings = randstr.String(16)
	// 	}
	// }

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello "+greetings+" !")
	})

	e.GET("/random", func(c echo.Context) error {
		return c.HTML(http.StatusOK, greetings)
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
