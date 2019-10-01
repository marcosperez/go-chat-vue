package handlers

import (
	"net/http"

	"github.com/marcosperez/go-chat-vue/models/common"

	"github.com/labstack/echo"
	"github.com/marcosperez/go-chat-vue/stores"
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
	// auth ??
	u := new(common.User)
	if err = c.Bind(u); err != nil {
		return
	}
	user, err := uh.stores.UsersStore.CreateUser(u.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "No se logro recuperar el listado de usuarios")
	}
	return c.JSON(http.StatusOK, user)
}

// GetUsers endpoint que retorna el listado de usuarios
func (uh UserHandler) GetUsers(c echo.Context) (err error) {
	// auth ??
	users, err := uh.stores.UsersStore.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "No se logro recuperar el listado de usuarios")
	}

	return c.JSON(http.StatusOK, users)
}
