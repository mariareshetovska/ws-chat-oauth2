package ws

import (
	"log"
	"net/http"
	"ws/controller"

	"github.com/gorilla/websocket"
)

type Server struct {
	upgrader websocket.Upgrader
	clients  []*websocket.Conn
}

func NewServer() *Server {
	return &Server{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		clients: []*websocket.Conn{},
	}
}

func (s *Server) EchoHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("websocket upgrade error: %s", err)
		return
	}
	defer conn.Close()

	s.clients = append(s.clients, conn)

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("websocket read error: %s", err)
			return
		}

		log.Printf("%s send: %s\n", conn.RemoteAddr(), string(msg))

		for _, client := range s.clients {
			if err = client.WriteMessage(msgType, msg); err != nil {
				log.Printf("websocket write error: %s", err)
				return
			}
		}
	}
}

func (s *Server) StaticHandler(w http.ResponseWriter, r *http.Request) {
	if controller.Session["code"] == "" {
		log.Printf("User is not authenticated. Redirecting to home page...")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	http.ServeFile(w, r, "chat.html")
}
