package socket

import (
	"time"

	"github.com/labstack/echo"
	"github.com/marcosperez/go-chat-vue/models/common"
	"github.com/rs/xid"
	"golang.org/x/net/websocket"
)

const timeoutKeepAlive = 60 * time.Second // TODO: terminar de definir

// Connection Mantiene el estado de una conexion web socket
type Connection struct {
	ID           string
	WS           *websocket.Conn
	timeout      time.Duration
	lastResponse time.Time
	logger       echo.Logger
	OnClose      func(conn *Connection)
}

// CreateConection Crea una conexion
func CreateConection(ws *websocket.Conn, l echo.Logger) *Connection {
	// Instanciacion y carga de datos por default
	conn := &Connection{
		ID:           xid.New().String(),
		WS:           ws,
		lastResponse: time.Now(),
		logger:       l,
	}
	conn.OnClose = func(conn *Connection) {}
	// Funcion que detecta el cierre de la conexion
	conn.startKeepAlive()

	return conn
}

// PongReceiver Actualiza fecha y hora de ultimo mensaje recibido
func (c *Connection) PongReceiver() {
	c.lastResponse = time.Now()
	c.logger.Debugf("\n[Connection][PongReceiver] Actualizacion de ultima respuesta %s", c.lastResponse.Format(time.RFC3339))
}

// startKeepAlive Funcion que detecta el cierre de la conexion
func (c *Connection) startKeepAlive() {
	go func() {
		defer func() {
			c.WS.Close()
			c.OnClose(c)
		}()

		for {
			// Enviar ping para verificar si esta activo el socket
			if err := websocket.JSON.Send(c.WS, common.SendMessage{Type: "ping"}); err != nil {
				c.logger.Debugf("\n[Connection][keepAlive] Error al enviar ping , error: %s", err)
			}
			time.Sleep(timeoutKeepAlive / 3)
			c.logger.Debugf("\n[Connection][keepAlive] Diferencia: %d", time.Now().Sub(c.lastResponse))
			if time.Now().Sub(c.lastResponse) > timeoutKeepAlive {
				c.logger.Debugf("\n[Connection][keepAlive] Cerrando conexion por KeepAlive")
				c.WS.Close()
				c.OnClose(c)
				return
			}
		}
	}()
}
