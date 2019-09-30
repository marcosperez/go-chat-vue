package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
)

func main() {
	server := configureSocket()
	// http.Handle("/ws", corsMiddleware(server))
	// http.Handle("/", http.FileServer(http.Dir("./web")))
	// log.Println("Serving at localhost:8357...")
	// http.ListenAndServe(":8357", nil)

	e := echo.New()

	e.Any("/ws", corsMiddleware(server))

	e.Any.Handle("/", http.FileServer(http.Dir("./web")))

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	e.Logger.Fatal(e.Start(":1323"))
}

func webSocketHandler(ws *websocket.Conn) {
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

func configureSocket() http.Handler {
	return websocket.Handler(webSocketHandler)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		next.ServeHTTP(w, r)
	})
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	// (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
