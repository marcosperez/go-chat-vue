package stores

import (
	"errors"

	"github.com/marcosperez/go-chat-vue/models/common"
	"github.com/rs/xid"
)

// UsersStore interfacaz de metodos para Store de usuarios
type UsersStore interface {
	CreateUser(name string) (*common.User, error)
	GetUsers() ([]common.User, error)
	GetUser(name string) (*common.User, error)
}

// UsersStoreMemory tipo de dato que implementa UsersStore
type UsersStoreMemory struct {
	users    []common.User          // base de datos en memoria ""
	usersMap map[string]common.User // Base de datos para acceso mas directo
}

// CreateUsersStoreMemory crea el store
func CreateUsersStoreMemory() *UsersStoreMemory {
	return &UsersStoreMemory{users: []common.User{}}
}

// CreateUser agrega un usuario a la "BD"
func (u *UsersStoreMemory) CreateUser(name string) (user *common.User, err error) {
	user = u.findUserByName(name)
	if user != nil {
		return user, nil
	}

	userID := xid.New()
	newUser := common.User{ID: userID.String(), Name: name}
	u.users = append(u.users, newUser)
	u.usersMap[userID.String()] = newUser
	return &newUser, nil
}

// GetUsers metodo para obtener el listado de usuarios
func (u *UsersStoreMemory) GetUsers() ([]common.User, error) {
	return u.users, nil
}

// GetUser metodo para obtener un usuario especifico por id
func (u *UsersStoreMemory) GetUser(id string) (*common.User, error) {
	user, exist := u.usersMap[id]
	if exist {
		return &user, nil
	}
	return nil, errors.New("No existe el usuario")
}

// GetUserByName metodo para obtener un usuario especifico por nombre
func (u *UsersStoreMemory) GetUserByName(name string) (*common.User, error) {
	user := u.findUserByName(name)
	if user != nil {
		return user, nil
	}
	return nil, errors.New("No existe el usuario")
}

// Metodos privados
func (u *UsersStoreMemory) findUserByName(name string) *common.User {

	for _, user := range u.users {
		if user.Name == name {
			return &user
		}
	}
	return nil
}
