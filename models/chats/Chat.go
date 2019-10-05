package chats

import (
	"time"

	"github.com/marcosperez/go-chat-vue/models/common"
	"github.com/rs/xid"
	"golang.org/x/net/websocket"
)

// ChatMessage generados por los usuarios
type ChatMessage struct {
	messageText string
	date        time.Time
}

// Chat Objeto que contiene un chat particular
type Chat struct {
	ID       string
	users    map[string]string
	messages []ChatMessage
	Channel  chan common.ChatMessage

	supervisor *Supervisor
}

// CreateChat crea un chat nuevo por un usuario
func CreateChat(chatID string, supervisor *Supervisor) (*Chat, error) {
	users := make(map[string]string)
	channel := make(chan common.ChatMessage)
	messages := []ChatMessage{}
	// Chat con ID bien conocido
	if chatID == "" {
		chat := Chat{ID: chatID, users: users, messages: messages, Channel: channel, supervisor: supervisor}
		chat.StartChat()
		return &chat, nil
	}
	// TODO mejorar implementacion
	// Chat random
	chat := &Chat{ID: xid.New().String(), users: users, messages: messages, Channel: channel, supervisor: supervisor}
	chat.StartChat()
	return chat, nil
}

// StartChat Inicia loop de chat
func (c *Chat) StartChat() {

	go func() {
		defer close(c.Channel)

		for msg := range c.Channel {
			for _, userID := range c.users {
				// El mismo usuario
				if userID == msg.UserID {
					break
				}

				conn, err := c.supervisor.ws.GetConnection((userID))
				if err != nil {
					c.supervisor.logger.Errorf("No existe conexion para el userID %s", userID)
					return
				}
				websocket.JSON.Send(conn.WS, common.SendMessage{Type: "chat", Data: msg})
			}
		}
	}()
}

// Join se une un nuevo usario al chat
func (c *Chat) Join(userID string) error {
	c.users[userID] = userID
	return nil
}
