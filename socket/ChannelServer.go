package socket

import (
	"github.com/marcosperez/go-chat-vue/models"
	"github.com/rs/xid"
	"golang.org/x/net/websocket"
)

// Channels por los cuales se comunican con el supervisor de channels
type ChannelsServer struct {
	MessagesChannel     chan models.ChannelMessage
	SuscriptionsChannel chan models.SubscriptionMessage
}

// ChannelServer Gestiona un channel
type ChannelServer struct {
	ID           string
	supervisor   *ChannelsSupervisor
	Channel      *models.Channel
	ChannelQueue chan models.ChannelMessage
}

// CreateChannelServer crea un channel nuevo por un usuario
func CreateChannelServer(supervisor *ChannelsSupervisor) (*ChannelServer, error) {
	users := make(map[string]string)
	channelQueue := make(chan models.ChannelMessage)
	messages := []models.ChannelMessage{}
	channel := &models.Channel{Users: users, Messages: messages} // Generar de otra forma

	// // Channel con ID bien conocido
	// if channelID == "" {
	// 	server := ChannelServer{ChannelQueue: channelQueue, Channel: channel, supervisor: supervisor}
	// 	server.StartChannel()
	// 	return &server, nil
	// }
	// TODO mejorar implementacion
	// Channel random
	c := &ChannelServer{ID: xid.New().String(), ChannelQueue: channelQueue, Channel: channel, supervisor: supervisor}
	c.StartChannel()
	return c, nil
}

// StartChannel Inicia loop de channel
func (c *ChannelServer) StartChannel() {

	go func() {
		defer close(c.ChannelQueue)

		for msg := range c.ChannelQueue {
			for _, userID := range c.Channel.Users {
				// El mismo usuario
				if userID == msg.UserID {
					break
				}

				conn, err := c.supervisor.ws.GetConnection((userID))
				if err != nil {
					c.supervisor.logger.Errorf("No existe conexion para el userID %s", userID)
					return
				}
				websocket.JSON.Send(conn.WS, models.SendMessage{Type: "channel", Data: msg})
			}
		}
	}()
}

// Join se une un nuevo usario al channel
func (c *ChannelServer) Join(userID string) error {
	c.Channel.Users[userID] = userID
	return nil
}
