package shoe

import (
	"net/http"

	"github.com/MaurickThom/Taller-APIs-REST-Golang/response_model"

	"github.com/labstack/echo"
)

func Create(context echo.Context) error {
	model := &Model{} // puntero del model
	err := context.Bind(model)
	if err != nil {
		resp := response_model.Model{
			MessageError: response_model.MessageError{
				"E102",
				"the shoe is badly sent",
			},
		}

		return context.JSON(http.StatusBadRequest, resp)
	}
	data := storage.Create(model)
	resp := response_model.Model{
		MessageOK: response_model.MessageOK{
			"A001",
			"shoe created correctly",
		},
		Data: data,
	}
	return context.JSON(http.StatusCreated, resp)
}
