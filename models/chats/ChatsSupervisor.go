package chats

import (
	"github.com/labstack/echo"
	"github.com/marcosperez/go-chat-vue/models/common"
	"github.com/marcosperez/go-chat-vue/socket"
	"github.com/marcosperez/go-chat-vue/stores"
)

// Supervisor encargado de gestion gorutinas de chat
type Supervisor struct {
	chats    map[string]*Chat
	logger   echo.Logger
	stores   *stores.Stores
	ws       *socket.Server
	Channels *common.Channels
}

// CreateChatsSupervisor crea un supervisor para un grupo de chats
func CreateChatsSupervisor() *Supervisor {
	channels := &common.Channels{
		ChatsChannel:        make(chan common.ChatMessage),
		SuscriptionsChannel: make(chan common.SubscriptionMessage),
	}
	return &Supervisor{
		chats:    make(map[string]*Chat),
		Channels: channels,
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
		for ChatMessage := range cs.Channels.ChatsChannel {
			cs.logger.Infof("[ChatSupervisor] Mensaje: %v", ChatMessage)
			chat := cs.GetChat(ChatMessage.ChatID)
			chat.Channel <- ChatMessage
		}
	}()
}

// StopChatSupervisor compeltar
func (cs *Supervisor) StopChatSupervisor() {
	close(cs.Channels.ChatsChannel)
	// completar
}
