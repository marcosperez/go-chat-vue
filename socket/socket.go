package socket

import (
	"errors"
	"fmt"
	"runtime/debug"

	"github.com/labstack/echo"
	"github.com/marcosperez/go-chat-vue/models"
	"github.com/marcosperez/go-chat-vue/stores"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/net/websocket"
)

// Server encargado de gestionar las conexiones WS
type Server struct {
	stores             *stores.Stores
	connections        map[string]*Connection
	logger             echo.Logger
	ChannelsSupervisor *ChannelsSupervisor
}

// CreateSocketServer Crea una instancia de SocketServer
func CreateSocketServer() *Server {
	cs := CreateChannelsSupervisor()

	return &Server{
		connections:        make(map[string]*Connection),
		ChannelsSupervisor: cs,
	}
}

// InjectDependencies metodo para inyectar dependencias al servidor de web socket
func (ss *Server) InjectDependencies(stores *stores.Stores, l echo.Logger) {
	ss.stores = stores

	ss.logger = l
	ss.ChannelsSupervisor.InjectDependencies(l, stores, ss)
}

// GetConnection retorna una conexion asociada a un usuario
func (ss *Server) GetConnection(userID string) (*Connection, error) {
	conn, ok := ss.connections[userID]
	if ok {
		return conn, nil
	}
	return nil, errors.New("La conection no existe")
}

// SocketHandler handler que recibe mensajes web socket
func (ss *Server) SocketHandler(c echo.Context) error {
	ss.ChannelsSupervisor.StartChannelsChannelsSupervisor()
	websocket.Handler(func(ws *websocket.Conn) {
		defer func() {
			ss.recoverWS()
			ws.Close()
		}()
		var err error

		for {
			var message models.ReceiverMessage

			if err = websocket.JSON.Receive(ws, &message); err != nil {
				c.Logger().Errorf("[WebSocket] Error al recibir el mensaje ")
				return
			}
			// Discriminador de tipo de evento/mensaje
			switch message.Type { // TODO: Pasar tipos a constantes
			case "connection":
				ss.onConnection(message, ws)

			case "suscription":
				ss.onSuscription(message)

			case "channel":
				ss.onReceiveChannel(message)

			case "pong":
				ss.onReceivePong(message)

			default:
				c.Logger().Debugf("[WebSocket] Mensaje no controlado %v ", message)
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func (ss *Server) onSuscription(message models.ReceiverMessage) {
	ss.logger.Debugf("[WebSocket] onSuscription %v ", message)
	var ChannelMessage models.SubscriptionMessage
	err := mapstructure.Decode(message.Data, &ChannelMessage)
	if err != nil {
		ss.logger.Errorf("[WebSocket] Error al convertir data en tipo  %v", err)
		return
	}
	ChannelMessage.UserID = message.UserID
	ss.logger.Infof("[WebSocket] \nChannelID: %v Mensaje: %v", ChannelMessage.ChannelID, ChannelMessage.UserID)
	// suscribir a un canal especifico
	ss.ChannelsSupervisor.Channels.SuscriptionsChannel <- ChannelMessage
}

func (ss *Server) onConnection(message models.ReceiverMessage, ws *websocket.Conn) {
	ss.logger.Debugf("[WebSocket] connection %v ", message)
	conn := CreateConection(ws, ss.logger)
	conn.OnClose = func(c *Connection) { delete(ss.connections, message.UserID) } // TODO mejorar mas elegantee
	ss.connections[message.UserID] = conn
}

func (ss *Server) onReceivePong(message models.ReceiverMessage) {
	ss.logger.Debugf("[WebSocket] Pong %v ", message)
	conn, exist := ss.connections[message.UserID]
	if exist {
		conn.PongReceiver()
	}
}

func (ss *Server) onReceiveChannel(message models.ReceiverMessage) {
	var ChannelMessage models.ChannelMessage
	err := mapstructure.Decode(message.Data, &ChannelMessage)
	if err != nil {
		ss.logger.Errorf("[WebSocket] Error al convertir data en tipo  %v", err)
		return
	}
	ss.logger.Infof("[WebSocket] \nChannelID: %v Mensaje: %v", ChannelMessage.ChannelID, ChannelMessage.Message)
	ChannelMessage.UserID = message.UserID
	// Enviar mensaje a al channel
	ss.ChannelsSupervisor.Channels.MessagesChannel <- ChannelMessage
}

func (ss *Server) recoverWS() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
		debug.PrintStack()
	}
}
