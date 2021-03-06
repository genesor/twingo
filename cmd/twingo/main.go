package main

import (
	"github.com/labstack/echo"

	"github.com/genesor/twingo/http"
)

func main() {
	e := echo.New()

	helloHandler := http.NewHelloHandler()
	e.GET("/", helloHandler.HandleHello)
	e.Logger.Fatal(e.Start(":9494"))
}
