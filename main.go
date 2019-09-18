package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", handlerHelloWorld)
	err := e.Start(":8080")
	if err != nil { // nil significa nulo , entonces se le pregunta si error
		// esta nulo o no, si no esta nulo significa que hay un error
		fmt.Printf("No pude subir el servidor %v", err)
	}
}

func handlerHelloWorld(context echo.Context) error {
	return context.String(http.StatusOK, "Hello world")
}
