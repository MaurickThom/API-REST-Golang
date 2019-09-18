package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	var e = echo.New()
	e.GET("/", handlerHelloWorld)
	var err = e.Start(":8080")
	if err != nil { // nil significa nulo , entonces se le pregunta si error
		// esta nulo o no, si no esta nulo significa que hay un error
		fmt.Printf("No pude subir el servidor %v", err)
	}
}

/*
con el framework echo necesito el contexto
en el contexto me trae el request y el response
*/
func handlerHelloWorld(context echo.Context) error {
	return context.String(http.StatusOK, "Hello world")
}
