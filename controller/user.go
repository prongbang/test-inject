package controller

import (
	"fmt"
	"net/http"
	"test-inject/model"

	"github.com/labstack/echo"
)

type Controller struct {
	Con *model.Connecttion
}

func (ctrl *Controller) Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func (ctrl *Controller) UserLogin(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	sql := "SELECT id, first_name, last_name, username, password, types FROM users WHERE username = '" + u.Username + "' AND password = '" + u.Password + "'"
	fmt.Println(sql)
	row, err := ctrl.Con.Db.Query(sql)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	obj := model.User{}
	if row.Next() {
		e := row.Scan(&obj.ID, &obj.FirstName, &obj.LastName, &obj.Username, &obj.Password, &obj.Types)
		if e != nil {
			fmt.Println(e)
			return c.JSON(http.StatusUnauthorized, e)
		}
	}
	if obj.ID == 0 {
		return c.JSON(http.StatusUnauthorized, "401 Unauthorized")
	}
	return c.JSON(http.StatusOK, obj)
}

func (ctrl *Controller) SaveUser(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func (ctrl *Controller) GetUserAll(c echo.Context) error {
	rows, err := ctrl.Con.Db.Query("SELECT id, first_name, last_name, username, password, types FROM users")
	list := []model.User{}
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	for rows.Next() {
		obj := model.User{}
		e := rows.Scan(&obj.ID, &obj.FirstName, &obj.LastName, &obj.Username, &obj.Password, &obj.Types)
		if e != nil {
			return c.JSON(http.StatusNotFound, e)
		}
		list = append(list, obj)
	}
	return c.JSON(http.StatusOK, list)
}

func (ctrl *Controller) GetUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusNotFound, id)
	}

	sql := "SELECT id, first_name, last_name, username, password, types FROM users WHERE id = " + id
	row, err := ctrl.Con.Db.Query(sql)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	obj := model.User{}
	if row.Next() {
		e := row.Scan(&obj.ID, &obj.FirstName, &obj.LastName, &obj.Username, &obj.Password, &obj.Types)
		if e != nil {
			return c.JSON(http.StatusNotFound, e)
		}
	}

	return c.JSON(http.StatusOK, obj)
}
