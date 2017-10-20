package main

import (
	"test-inject/controller"
	"test-inject/gdbc"
	"test-inject/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func main() {

	db := gdbc.GetConnection()
	con := &model.Connecttion{db}
	ctrl := controller.Controller{con}

	e := echo.New()
	e.GET("/", ctrl.Home)
	e.POST("/users", ctrl.SaveUser)
	e.GET("/users", ctrl.GetUserAll)
	e.GET("/users/:id", ctrl.GetUser)
	e.POST("/users/login", ctrl.UserLogin)

	e.Logger.Fatal(e.Start(":9000"))
}
