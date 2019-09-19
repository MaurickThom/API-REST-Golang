package main

import (
	"github.com/MaurickThom/Taller-APIs-REST-Golang/shoe"
	"github.com/MaurickThom/Taller-APIs-REST-Golang/user"
	"github.com/labstack/echo"
)

func startRoutes(e *echo.Echo) {
	e.POST("/shoes", shoe.Create, user.ValidateJWT)
	e.POST("/users", user.Login)

	// para ver todos los usuarios se debe de ser del role de admin
	// con un simple middleware , como no domino muy bien go lo hare sin jwt
	// TODO : middleware jwt -- ROLE: ADMIN
	e.GET("/users", user.GetAll)

}
