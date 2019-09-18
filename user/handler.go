package user

import (
	"net/http"

	"github.com/MaurickThom/Taller-APIs-REST-Golang/response_model"
	"github.com/labstack/echo"
)

func Login(context echo.Context) error {
	user := &Model{}
	err := context.Bind(user)
	if err != nil {
		resp := response_model.Model{
			MessageError: response_model.MessageError{
				Code:    "E003",
				Content: "badly sent object",
			},
		}
		return context.JSON(http.StatusBadRequest, resp)
	}
	dbUser := storage.Login(user.Email, user.Password)
	if dbUser == nil {
		resp := response_model.Model{
			MessageError: response_model.MessageError{
				Code:    "L001",
				Content: "Incorrect password or email",
			},
		}
		return context.JSON(http.StatusBadRequest, resp)
	}
	return context.JSON(http.StatusOK, dbUser)
}
