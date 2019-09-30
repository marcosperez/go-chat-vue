package chats

import (
	"time"

	"github.com/rs/xid"
)

// Message generados por los usuarios
type Message struct {
	messageText string
	date        time.Time
}

// Chat Objeto que contiene un chat particular
type Chat struct {
	ID       string
	users    []string
	messages []Message
}

// CreateChat crea un chat nuevo por un usuario
func CreateChat(userID string) (Chat, error) {
	users := []string{userID}
	messages := []Message{}
	chatID := xid.New()

	chat := Chat{ID: chatID.String(), users: users, messages: messages}
	return chat, nil
}

// Join se une un nuevo usario al chat
func (c *Chat) Join(userID string) error {
	c.users = append(c.users, userID)
	return nil
}
