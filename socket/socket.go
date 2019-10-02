package socket

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/echo"
	"github.com/marcosperez/go-chat-vue/models/chats"
	"golang.org/x/net/websocket"
)

// Server encargado de gestionar las conexiones WS
type Server struct {
	chatsSupervisor *chats.ChatsSupervisor
}

// CreateSocketServer Crea una instancia de SocketServer
func CreateSocketServer() *Server {
	return &Server{}
}

// InjectDependencies metodo para inyectar dependencias al servidor de web socket
func (ss *Server) InjectDependencies(sc *chats.ChatsSupervisor) {
	ss.chatsSupervisor = sc
}

// SocketHandler handler que recibe mensajes web socket
func (ss *Server) SocketHandler(c echo.Context) error {
	websocket.Handler(ss.createSocketHandler).ServeHTTP(c.Response(), c.Request())
	return nil
}

func (ss *Server) createSocketHandler(ws *websocket.Conn) {
	defer ws.Close()
	var err error

	for {
		var messageString string

		if err = websocket.Message.Receive(ws, &messageString); err != nil {
			fmt.Println("[WebSocket] Error al recibir el mensaje ")
			break
		}
		// Parseo de mensaje recibido
		var message Message
		err = json.Unmarshal([]byte(messageString), &message)
		if err != nil {
			fmt.Println("[WebSocket] Error al deserializar mensaje", messageString)
		}

		// Discriminador de tipo de evento/mensaje
		switch message.Type {
		case "subscripcion":
			fmt.Printf("[WebSocket] subscripcion %v ", message)
			ss.chatsSupervisor.Channel <- message
		default:
			ss.chatsSupervisor.Channel <- message
			fmt.Printf("[WebSocket] Mensaje no controlado %v ", message)
		}

		// if err = websocket.Message.Send(ws, msg); err != nil {
		// 	fmt.Println("Can't send")
		// 	break
		// }
	}
}

// Message datos recibidos por WS
type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

// SubscriptionData data de un mensaje de tipo subscripcion
type SubscriptionData struct {
	Room string `json:"room"`
}
