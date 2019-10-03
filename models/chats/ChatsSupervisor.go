package chats

import (
	"github.com/labstack/echo"
	"github.com/marcosperez/go-chat-vue/models/common"
	"github.com/marcosperez/go-chat-vue/socket"
	"github.com/marcosperez/go-chat-vue/stores"
)

// Supervisor encargado de gestion gorutinas de chat
type Supervisor struct {
	chats  map[string]*Chat
	logger echo.Logger
	stores *stores.Stores
	ws     *socket.Server

	Channel chan common.ChatData
}

// CreateChatsSupervisor crea un supervisor para un grupo de chats
func CreateChatsSupervisor() *Supervisor {
	return &Supervisor{
		chats:   make(map[string]*Chat),
		Channel: make(chan common.ChatData),
	}
}

// InjectDependencies metodo para inyectar dependencias
func (cs *Supervisor) InjectDependencies(l echo.Logger, s *stores.Stores, ws *socket.Server) {
	cs.logger = l
	cs.stores = s
	cs.ws = ws
}

// SuscribeUser Suscribe a un usuario a todos sus chats
func (cs *Supervisor) SuscribeUser(u *common.User) {
	if u == nil {
		return
	}
	for _, chatID := range u.ChatIDs {
		chat := cs.GetChat(chatID)
		chat.Join(u.ID)
	}
}

// GetChat Obtiene o crea un chat por ID
func (cs *Supervisor) GetChat(chatID string) *Chat {
	chat, ok := cs.chats[chatID]
	if ok {
		return chat
	}

	chat, err := CreateChat(chatID, cs)
	if err != nil {
		cs.logger.Infof("Creando chat ID: %v", chatID)
	}
	cs.chats[chatID] = chat
	return chat
}

// StartChatSupervisor inicia el loop de receopcion de mensajes
func (cs *Supervisor) StartChatSupervisor() {

	go func() {
		defer func() {
			cs.StopChatSupervisor()
		}()
		// Loop infinito de mensajeria
		for chatData := range cs.Channel {
			cs.logger.Infof("[ChatSupervisor] Mensaje: %v", chatData)
			chat := cs.GetChat(chatData.ChatID)
			chat.Channel <- chatData
		}
	}()
}

// StopChatSupervisor compeltar
func (cs *Supervisor) StopChatSupervisor() {
	close(cs.Channel)
	// completar
}
