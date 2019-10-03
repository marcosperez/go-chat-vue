package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/marcosperez/go-chat-vue/handlers"
	"github.com/marcosperez/go-chat-vue/models/chats"
	"github.com/marcosperez/go-chat-vue/socket"
	"github.com/marcosperez/go-chat-vue/stores"
)

func main() {
	e := echo.New()
	e.Logger.SetHeader("[${time_rfc3339}][${level}][${short_file}:${line}]")
	e.Logger.SetLevel(2)
	e.Use(middleware.Recover())
	e.Use(corsMiddleware())
	// Instanciacion de stores (y services)
	stores := stores.InitStores()
	// InitSupervisor de chat
	chatsSupervisor := chats.CreateChatsSupervisor()
	// Archivos estaticos
	e.Static("/", "./web")
	// Web socket
	socketServer := socket.CreateSocketServer()
	e.GET("/ws", socketServer.SocketHandler)
	// Configuracion de api
	apiConfiguration(e, stores, chatsSupervisor)

	// Inyeccion de dependencias
	chatsSupervisor.InjectDependencies(e.Logger, stores, socketServer)
	socketServer.InjectDependencies(stores, chatsSupervisor.Channel)

	// Start server
	chatsSupervisor.StartChatSupervisor()
	e.Logger.Fatal(e.Start(":8357"))
}

func apiConfiguration(e *echo.Echo, stores *stores.Stores, cs *chats.Supervisor) {
	// Definicion de api
	// USERS
	g := e.Group("/api")
	UserHandler := handlers.CreateUserHandler(stores, cs)
	g.POST("/users", UserHandler.CreateUser)
	g.GET("/users", UserHandler.GetUsers)
}

// TODO: Mover a otro archivo
// func wsHandler(c echo.Context) error {
// 	websocket.Handler(webSocketHandler).ServeHTTP(c.Response(), c.Request())
// 	return nil
// }

// func webSocketHandler(ws *websocket.Conn) {
// 	defer ws.Close()
// 	var err error

// 	for {
// 		var reply string

// 		if err = websocket.Message.Receive(ws, &reply); err != nil {
// 			fmt.Println("Can't receive")
// 			break
// 		}

// 		fmt.Println("Received back from client: " + reply)

// 		msg := "Received:  " + reply
// 		fmt.Println("Sending to client: " + msg)

// 		if err = websocket.Message.Send(ws, msg); err != nil {
// 			fmt.Println("Can't send")
// 			break
// 		}
// 	}
// }

func corsMiddleware() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://127.0.0.1:5500"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	})
}
