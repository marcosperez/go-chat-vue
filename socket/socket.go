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

	chanMessages chan common.ChatData
}

// CreateSocketServer Crea una instancia de SocketServer
func CreateSocketServer() *Server {
	return &Server{connections: make(map[string]*Connection)}
}

// InjectDependencies metodo para inyectar dependencias al servidor de web socket
func (ss *Server) InjectDependencies(stores *stores.Stores, c chan common.ChatData) {
	ss.stores = stores
	ss.chanMessages = c
}

func (ss *Server) GetConnection(userID string) (*Connection, error) {
	conn, ok := ss.connections[userID]
	if ok {
		return conn, nil
	}
	return nil, errors.New("La conection no existe")
}

// SocketHandler handler que recibe mensajes web socket
func (ss *Server) SocketHandler(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ss.recoverWS()
		defer ws.Close()
		var err error

		for {
			var message ReceiverMessage

			if err = websocket.JSON.Receive(ws, &message); err != nil {
				c.Logger().Errorf("[WebSocket] Error al recibir el mensaje ")
				return
			}
			// Discriminador de tipo de evento/mensaje
			switch message.Type { // TODO: Pasar tipos a constantes
			case "connection":
				c.Logger().Debugf("[WebSocket] connection %v ", message)
				conn := CreateConection(ws, c.Logger())
				conn.OnClose = func(c *Connection) { delete(ss.connections, message.UserID) } // TODO mejorar mas elegantee
				ss.connections[message.UserID] = conn

			case "subscripcion":
				c.Logger().Debugf("[WebSocket] subscripcion %v ", message)

			case "chat":
				var chatData common.ChatData
				err := mapstructure.Decode(message.Data, &chatData)
				if err != nil {
					c.Logger().Errorf("[WebSocket] Error al convertir data en tipo  %v", err)
					return
				}
				c.Logger().Infof("[WebSocket] \nChatID: %v Mensaje: %v", chatData.ChatID, chatData.Message)
				chatData.UserID = message.UserID
				// ss.chatsSupervisor.OnMessage(chatData.ChatID, chatData.Message)
				ss.chanMessages <- chatData

			case "pong":
				c.Logger().Debugf("[WebSocket] Pong %v ", message)
				conn, exist := ss.connections[message.UserID]
				if exist {
					conn.PongReceiver()
				}

			default:
				c.Logger().Debugf("[WebSocket] Mensaje no controlado %v ", message)
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func (ss *Server) recoverWS() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
		debug.PrintStack()
	}
}

// ReceiverMessage datos recibidos por WS
type ReceiverMessage struct {
	Type   string      `json:"type"`
	UserID string      `json:"userID"` // TODO: Buscar forma mas elegante de identificar una conexion
	Data   interface{} `json:"data"`
}

// SubscriptionData data de un mensaje de tipo subscripcion
type SubscriptionData struct {
	ChatID string `json:"chatID"`
}
