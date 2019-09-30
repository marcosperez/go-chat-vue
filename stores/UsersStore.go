package stores

import (
	"errors"

	"../models/common"
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
	users []common.User // base de datos en memoria ""
}

// CreateUsersStoreMemory crea el store
func CreateUsersStoreMemory() *UsersStoreMemory {
	return &UsersStoreMemory{users: []common.User{}}
}

// CreateUser agrega un usuario a la "BD"
func (u *UsersStoreMemory) CreateUser(name string) (user *common.User, err error) {
	user = u.findUser(name)
	if user != nil {
		return user, nil
	}

	userID := xid.New()
	newUser := common.User{ID: userID.String(), Name: name}
	u.users = append(u.users, newUser)
	return &newUser, nil
}

// GetUsers metodo para obtener el listado de usuarios
func (u *UsersStoreMemory) GetUsers() ([]common.User, error) {
	return u.users, nil
}

// GetUser metodo para obtener un usuario especifico por nombre (podria crearse una estructura de filtro)
func (u *UsersStoreMemory) GetUser(name string) (*common.User, error) {
	user := u.findUser(name)
	if user != nil {
		return user, nil
	}
	return nil, errors.New("No existe el usuario")
}

func (u *UsersStoreMemory) findUser(name string) *common.User {

	for _, user := range u.users {
		if user.Name == name {
			return &user
		}
	}
	return nil
}
