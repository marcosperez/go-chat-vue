package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/marcosperez/go-chat-vue/stores"
)

// AuthHandler capturador de request de usuario
type AuthHandler struct {
	stores *stores.Stores
	// channelsSupervisor *channels.Supervisor
}

// CreateAuthHandler instanciacion de handler
func CreateAuthHandler(stores *stores.Stores) AuthHandler {
	return AuthHandler{stores: stores}
}

// Login login de usuario
func (ah AuthHandler) Login(c echo.Context) (err error) {

	return nil
}

// RequestSignup request body para registro
type RequestSignup struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Signup registro de usuario
func (ah AuthHandler) Signup(c echo.Context) (err error) {
	req := new(RequestSignup)
	if err = c.Bind(req); err != nil {
		return
	}

	u, err := ah.stores.UsersStore.CreateUser(req.Name, req.Password, req.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Error en registro de usuario")
	}

	return c.JSON(http.StatusOK, u)
}
