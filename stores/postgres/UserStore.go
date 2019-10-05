package postgres

import (
	"errors"

	"github.com/go-pg/pg"
	"github.com/marcosperez/go-chat-vue/models"
)

// UsersStorePG tipo de dato que implementa UsersStore
type UsersStorePG struct {
	db *pg.DB
}

// CreateUsersStorePG crea el store
func CreateUsersStorePG() *UsersStorePG {
	return &UsersStorePG{}
}

// CreateUser agrega un usuario a la "BD"
func (u *UsersStorePG) CreateUser(name string, password string, email string) (user *models.User, err error) {
	db := CreateDbConnection()
	user = &models.User{Name: name, Email: email, Password: password}

	err = db.Insert(user)
	if err != nil {
		return nil, errors.New("Error al insertar el usuario")
	}
	return user, nil
}

// GetUsers metodo para obtener el listado de usuarios
func (u *UsersStorePG) GetUsers() ([]models.User, error) {
	return nil, errors.New("No implementado")
}

// GetUser metodo para obtener un usuario especifico por id
func (u *UsersStorePG) GetUser(id string) (*models.User, error) {
	return nil, errors.New("No implementado")
}

// GetUserByName metodo para obtener un usuario especifico por nombre
func (u *UsersStorePG) GetUserByName(name string) (*models.User, error) {
	return nil, errors.New("No implementado")
}
