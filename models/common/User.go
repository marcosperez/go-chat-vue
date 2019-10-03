package common

// User clase basica de usuario
type User struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	ChatIDs []string `json:"-"`
}
