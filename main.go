package main

import (
	"fmt"

	"./handlers"
	"./stores"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/net/websocket"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(corsMiddleware())
	// Instanciacion de stores (y services)
	stores := stores.InitStores()
	// Archivos estaticos
	e.Static("/", "./web")
	// Web socket
	e.GET("/ws", wsHandler)
	// Configuracion de api
	apiConfiguration(e, stores)
	// Start server
	e.Logger.Fatal(e.Start(":8357"))
}

func apiConfiguration(e *echo.Echo, stores *stores.Stores) {
	// Definicion de api
	// USERS
	g := e.Group("/api")
	UserHandler := handlers.CreateUserHandler(stores)
	g.POST("/users", UserHandler.CreateUser)
}

// TODO: Mover a otro archivo
func wsHandler(c echo.Context) error {
	websocket.Handler(webSocketHandler).ServeHTTP(c.Response(), c.Request())
	return nil
}

func webSocketHandler(ws *websocket.Conn) {
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

func corsMiddleware() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://labstack.com", "https://labstack.net"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	})
}
