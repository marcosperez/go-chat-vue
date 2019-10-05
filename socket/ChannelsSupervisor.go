package socket

import (
	"github.com/labstack/echo"
	"github.com/marcosperez/go-chat-vue/models"
	"github.com/marcosperez/go-chat-vue/stores"
)

// ChannelsSupervisor encargado de gestion gorutinas de channel
type ChannelsSupervisor struct {
	servers  map[string]*ChannelServer
	logger   echo.Logger
	stores   *stores.Stores
	ws       *Server
	Channels *ChannelsServer
}

// CreateChannelsSupervisor crea un supervisor para un grupo de channels
func CreateChannelsSupervisor() *ChannelsSupervisor {
	channels := &ChannelsServer{
		MessagesChannel:     make(chan models.ChannelMessage),
		SuscriptionsChannel: make(chan models.SubscriptionMessage),
	}
	return &ChannelsSupervisor{
		servers:  make(map[string]*ChannelServer),
		Channels: channels,
	}
}

// InjectDependencies metodo para inyectar dependencias
func (cs *ChannelsSupervisor) InjectDependencies(l echo.Logger, s *stores.Stores, ws *Server) {
	cs.logger = l
	cs.stores = s
	cs.ws = ws
}

// SuscribeUser Suscribe a un usuario a todos sus channels
func (cs *ChannelsSupervisor) SuscribeUser(userID string) {
	user, err := cs.stores.UsersStore.GetUser(userID)
	if err != nil {
		cs.logger.Infof("[SuscribeUser] Error al obtener usuario %v", err)
	}
	if user == nil {
		return
	}
	// for _, channel := range user.Channels {
	// 	channel := cs.GetChannel(channel.)
	// 	channel.Join(user.ID)
	// }
}

// GetChannel Obtiene o crea un channel por ID
func (cs *ChannelsSupervisor) GetChannel(channelID string) *ChannelServer {
	channel, ok := cs.servers[channelID]
	if ok {
		return channel
	}

	channel, err := CreateChannelServer(cs)
	if err != nil {
		cs.logger.Infof("[CreateChannel] Creando channel ID: %v", channelID)
	}
	cs.servers[channelID] = channel
	return channel
}

// StartChannelsChannelsSupervisor inicia el loop de receopcion de mensajes
func (cs *ChannelsSupervisor) StartChannelsChannelsSupervisor() {
	go cs.startChats()
	go cs.startSuscriptions()
}

func (cs *ChannelsSupervisor) startSuscriptions() {
	defer cs.StopChannelChannelsSupervisor()
	// Loop infinito de mensajeria
	for msg := range cs.Channels.SuscriptionsChannel {
		cs.logger.Infof("[ChannelChannelsSupervisor][Suscripcion] Mensaje: %v", msg)
		channel := cs.GetChannel(msg.ChannelID)
		channel.Join(msg.UserID)
	}
}

func (cs *ChannelsSupervisor) startChats() {
	defer cs.StopChannelChannelsSupervisor()
	// Loop infinito de mensajeria
	for msg := range cs.Channels.MessagesChannel {
		cs.logger.Infof("[ChannelChannelsSupervisor][Chat] Mensaje: %v", msg)
		channel := cs.GetChannel(msg.ChannelID)
		channel.ChannelQueue <- msg
	}
}

// StopChannelChannelsSupervisor compeltar
func (cs *ChannelsSupervisor) StopChannelChannelsSupervisor() {
	close(cs.Channels.MessagesChannel)
	// completar
}
