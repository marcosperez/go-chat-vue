package stores

// Stores agrupa todos los stores
type Stores struct {
	UsersStore UsersStore
}

func InitStores() *Stores {
	return &Stores{
		UsersStore: CreateUsersStoreMemory(),
	}
}
