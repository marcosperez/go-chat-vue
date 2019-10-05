package socket

import (
	"errors"
	"fmt"
	"runtime/debug"

	"github.com/labstack/echo"
	"github.com/marcosperez/go-chat-vue/models/common"
	"github.com/marcosperez/go-chat-vue/stores"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/net/websocket"
)

// Server encargado de gestionar las conexiones WS
type Server struct {
	stores      *stores.Stores
	connections map[string]*Connection
	channels    *common.Channels
	logger      echo.Logger
}

// CreateSocketServer Crea una instancia de SocketServer
func CreateSocketServer() *Server {
	return &Server{connections: make(map[string]*Connection)}
}

// InjectDependencies metodo para inyectar dependencias al servidor de web socket
func (ss *Server) InjectDependencies(stores *stores.Stores, channels *common.Channels) {
	ss.stores = stores
	ss.channels = channels
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
	ss.logger = c.Logger()
	websocket.Handler(func(ws *websocket.Conn) {
		defer func() {
			ss.recoverWS()
			ws.Close()
		}()
		var err error

		for {
			var message common.ReceiverMessage

			if err = websocket.JSON.Receive(ws, &message); err != nil {
				c.Logger().Errorf("[WebSocket] Error al recibir el mensaje ")
				return
			}
			// Discriminador de tipo de evento/mensaje
			switch message.Type { // TODO: Pasar tipos a constantes
			case "connection":
				ss.onConnection(message, ws)

			case "suscription":
				ss.suscription(message)

			case "chat":
				ss.receiveChat(message)

			case "pong":
				ss.receivePong(message)

			default:
				c.Logger().Debugf("[WebSocket] Mensaje no controlado %v ", message)
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func (ss *Server) suscription(message common.ReceiverMessage) {
	ss.logger.Debugf("[WebSocket] suscription %v ", message)
	var ChatMessage common.ChatMessage
	err := mapstructure.Decode(message.Data, &ChatMessage)
	if err != nil {
		ss.logger.Errorf("[WebSocket] Error al convertir data en tipo  %v", err)
		return
	}
	ss.logger.Infof("[WebSocket] \nChatID: %v Mensaje: %v", ChatMessage.ChatID, ChatMessage.Message)
	// suscribir a un canal especifico
	ss.channels.ChatsChannel <- ChatMessage
}

func (ss *Server) onConnection(message common.ReceiverMessage, ws *websocket.Conn) {
	ss.logger.Debugf("[WebSocket] connection %v ", message)
	conn := CreateConection(ws, ss.logger)
	conn.OnClose = func(c *Connection) { delete(ss.connections, message.UserID) } // TODO mejorar mas elegantee
	ss.connections[message.UserID] = conn
}

func (ss *Server) receivePong(message common.ReceiverMessage) {
	ss.logger.Debugf("[WebSocket] Pong %v ", message)
	conn, exist := ss.connections[message.UserID]
	if exist {
		conn.PongReceiver()
	}
}

func (ss *Server) receiveChat(message common.ReceiverMessage) {
	var ChatMessage common.ChatMessage
	err := mapstructure.Decode(message.Data, &ChatMessage)
	if err != nil {
		ss.logger.Errorf("[WebSocket] Error al convertir data en tipo  %v", err)
		return
	}
	ss.logger.Infof("[WebSocket] \nChatID: %v Mensaje: %v", ChatMessage.ChatID, ChatMessage.Message)
	ChatMessage.UserID = message.UserID
	// Enviar mensaje a al chat
	ss.channels.ChatsChannel <- ChatMessage
}

func (ss *Server) recoverWS() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
		debug.PrintStack()
	}
}
