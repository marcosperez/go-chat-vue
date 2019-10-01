package socket

import (
	"fmt"

	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
)

// Server encargado de gestionar las conexiones WS
type Server struct {
}

// CreateSocketServer Crea una instancia de SocketServer
func CreateSocketServer() *Server {
	return &Server{}
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
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}
