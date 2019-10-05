package stores

import (
	"github.com/marcosperez/go-chat-vue/stores/postgres"
)

// Stores agrupa todos los stores
type Stores struct {
	UsersStore UsersStore
}

// InitStores instancia todos los stores
func InitStores() *Stores {
	userStore := postgres.CreateUsersStorePG()
	userStore.CreateUser("user", "pass", "mail")
	return &Stores{
		// UsersStore: memory.CreateUsersStoreMemory(),
		UsersStore: userStore,
	}
}
