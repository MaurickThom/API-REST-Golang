package main

import (
	"github.com/MaurickThom/Taller-APIs-REST-Golang/shoe"
	"github.com/MaurickThom/Taller-APIs-REST-Golang/user"
	"github.com/labstack/echo"
)

func startRoutes(e *echo.Echo) {
	e.POST("/shoes", shoe.Create)
	e.POST("/users", user.Login)
}
