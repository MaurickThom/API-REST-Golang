package shoe

import (
	"net/http"

	"github.com/labstack/echo"
)

func Create(context echo.Context) error {
	model := &Model{} // puntero del model
	err := context.Bind(model)
	if err != nil {
		return context.JSON(http.StatusBadRequest, "the object sent is not correct")
	}
	storage.Create(model)
	return context.JSON(http.StatusCreated, "OK")
}
