package stores

// Stores agrupa todos los stores
type Stores struct {
	UsersStore UsersStore
}

// InitStores instancia todos los stores
func InitStores() *Stores {
	return &Stores{
		UsersStore: CreateUsersStoreMemory(),
	}
}
