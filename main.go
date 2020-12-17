package main

import (
	"github.com/labstack/echo/v4"
	"go.elastic.co/apm/module/apmechov4"
	"log"
)

func main() {
	e := echo.New()
	e.Use(apmechov4.Middleware())
	e.GET("/ping", ping)

	log.Fatal(e.Start(":8093"))
}

func ping(e echo.Context) error {
	return e.JSON(200, map[string]string{
		"message": "ok",
	})
}
