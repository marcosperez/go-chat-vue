package memory

import (
	"errors"
	"fmt"

	"github.com/marcosperez/go-chat-vue/models"
)

// UsersStoreMemory tipo de dato que implementa UsersStore
type UsersStoreMemory struct {
	users    []models.User          // base de datos en memoria ""
	usersMap map[string]models.User // Base de datos para acceso mas directo
}

// CreateUsersStoreMemory crea el store
func CreateUsersStoreMemory() *UsersStoreMemory {
	return &UsersStoreMemory{users: []models.User{}, usersMap: make(map[string]models.User)}
}

// CreateUser agrega un usuario a la "BD"
func (u *UsersStoreMemory) CreateUser(name string, password string, email string) (user *models.User, err error) {
	user, err = u.GetUserByName(name)
	if err != nil {
		fmt.Printf("Error de creacion de usuario, no se pudo obtener.")
	}
	if user != nil {
		return user, nil
	}

	userID := name //xid.New().String()
	newUser := models.User{
		ID:   userID,
		Name: name,
		// ChannelIDs: []string{"global", "users"},
	}
	// Base de datos en memoria
	u.users = append(u.users, newUser)
	u.usersMap[userID] = newUser
	return &newUser, nil
}

// GetUsers metodo para obtener el listado de usuarios
func (u *UsersStoreMemory) GetUsers() ([]models.User, error) {
	return u.users, nil
}

// GetUser metodo para obtener un usuario especifico por id
func (u *UsersStoreMemory) GetUser(id string) (*models.User, error) {
	user, exist := u.usersMap[id]
	if exist {
		return &user, nil
	}
	return nil, errors.New("No existe el usuario")
}

// GetUserByName metodo para obtener un usuario especifico por nombre
func (u *UsersStoreMemory) GetUserByName(name string) (*models.User, error) {
	for _, user := range u.users {
		if user.Name == name {
			return &user, nil
		}
	}
	return nil, errors.New("No existe el usuario")
}
