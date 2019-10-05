package models

// User clase basica de usuario
type User struct {
	tableName struct{}  `pg:"users"`
	ID        string    `json:"id" pg:"default:gen_random_uuid()"`
	Name      string    `json:"name" `
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Channels  []Channel `json:"-" pg:"many2many:channels"`
}
