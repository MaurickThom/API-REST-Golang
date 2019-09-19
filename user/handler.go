package user

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

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
	token, err := generateJWT(*dbUser)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "no se pudo generar el token")
	}
	type logueo struct {
		User  Model
		Token string
	}

	l := logueo{
		*dbUser,
		token,
	}

	resp := response_model.Model{
		MessageOK: response_model.MessageOK{
			Code:    "O001",
			Content: "Logueado ok",
		},
		Data: l,
	}

	return context.JSON(http.StatusOK, resp)
}

// getTokenFromAuthorizationHeader busca el token del header Authorization
func getTokenFromAuthorizationHeader(r *http.Request) (string, error) {
	ah := r.Header.Get("Authorization")
	if ah == "" {
		return "", errors.New("el encabezado no contiene la autorización")
	}

	// Should be a bearer token
	if len(ah) > 6 && strings.ToUpper(ah[0:6]) == "BEARER" {
		return ah[7:], nil
	} else {
		return "", errors.New("el header no contiene la palabra Bearer")
	}
}

// getTokenFromURLParams busca el token de la URL
func getTokenFromURLParams(r *http.Request) (string, error) {
	ah := r.URL.Query().Get("authorization")
	if ah == "" {
		return "", errors.New("la URL no contiene la autorización")
	}

	return ah, nil
}
func GetAll(context echo.Context) error {
	queryString := context.QueryParams()
	if len(queryString) == 0 {
		us := storage.GetAll()
		r := response_model.Model{
			MessageOK: response_model.MessageOK{
				Code:    "U205",
				Content: "Consultado correctamente",
			},
			Data: us,
		}
		return context.JSON(http.StatusOK, r)
	}

	return GetAllPaginate(context)
}
func GetAllPaginate(context echo.Context) error {
	l := context.QueryParam("limit")
	p := context.QueryParam("page")

	limit, err := strconv.Atoi(l)
	if err != nil {
		limit = 1
	}
	page, err := strconv.Atoi(p)
	if err != nil {
		page = 1
	}

	us := storage.GetAllPaginate(limit, page)

	r := response_model.Model{
		MessageOK: response_model.MessageOK{
			Code:    "U205",
			Content: "Consultado correctamente",
		},
		Data: us,
	}

	return context.JSON(http.StatusOK, r)
}
