package stores

import (
	"github.com/marcosperez/go-chat-vue/models"
)

// UsersStore interfacaz de metodos para Store de usuarios
type UsersStore interface {
	CreateUser(name string, password string, email string) (*models.User, error)
	GetUsers() ([]models.User, error)
	GetUser(name string) (*models.User, error)
	GetUserByName(name string) (*models.User, error)
}
