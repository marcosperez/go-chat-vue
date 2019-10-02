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
func CreateChat(chatID string) (Chat, error) {
	users := []string{}
	messages := []Message{}
	// Chat con ID bien conocido
	if chatID != "" {
		chat := Chat{ID: chatID, users: users, messages: messages}
		return chat, nil
	}
	// Chat random
	return Chat{ID: xid.New().String(), users: users, messages: messages}, nil
}

// Join se une un nuevo usario al chat
func (c *Chat) Join(userID string) error {
	c.users = append(c.users, userID)
	return nil
}
