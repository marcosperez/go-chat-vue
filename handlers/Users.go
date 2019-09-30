package handlers

import (
	"net/http"

	"../models/common"

	"../stores"
	"github.com/labstack/echo"
)

// UserHandler capturador de request de usuario
type UserHandler struct {
	stores *stores.Stores
}

// CreateUserHandler instanciacion de handler
func CreateUserHandler(stores *stores.Stores) UserHandler {
	return UserHandler{stores: stores}
}

// CreateUser endpoint de creacion de usuarios
func (uh UserHandler) CreateUser(c echo.Context) (err error) {

	u := new(common.User)
	if err = c.Bind(u); err != nil {
		return
	}

	return c.JSON(http.StatusOK, u)
}
