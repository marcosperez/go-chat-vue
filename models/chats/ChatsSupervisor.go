package chats

import (
	"errors"
	"fmt"
)

// ChatsSupervisor encargado de gestion gorutinas de chat
type ChatsSupervisor struct {
	chats   []Chat
	Channel chan interface{}
}

// CreateChatsSupervisor crea un supervisor para un grupo de chats
func CreateChatsSupervisor() *ChatsSupervisor {
	return &ChatsSupervisor{}
}

// InitChatsSupervisor inicializacion de valores requeredios
func (cs *ChatsSupervisor) InitChatsSupervisor() error {
	var global, err = CreateChat("global")
	if err != nil {
		return errors.New("Error en inicializacion de ChatsSupervisor")
	}

	// Siempre va a existir un chat con ID global
	cs.chats = []Chat{global}

	// creacion de channel
	cs.Channel = make(chan interface{})
	return nil
}

// StartChatSupervisor inicia el loop de receopcion de mensajes
func (cs *ChatsSupervisor) StartChatSupervisor() {
	defer close(cs.Channel)
	// Loop infinito de mensajeria
	for message := range cs.Channel {
		fmt.Printf("[ChatSupervisor] Mensaje: %v", message)
	}
}
