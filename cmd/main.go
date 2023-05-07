package main

import (
	"log"
	"net/http"
	"ws/controller"
	ws "ws/ws"
)

func main() {
	s := ws.NewServer()

	http.HandleFunc("/", controller.HandleMain)
	http.HandleFunc("/auth/google/login", controller.HandleGoogleLogin)
	http.HandleFunc("/auth/google/callback", controller.HandleGoogleCallback)

	http.HandleFunc("/auth/linkedin/login", controller.HandleLinkedinLogin)
	http.HandleFunc("/auth/linkedin/callback", controller.HandleLinkedinCallback)

	http.HandleFunc("/echo", s.EchoHandler)
	http.HandleFunc("/chat", s.StaticHandler)

	log.Println("Server started...")
	err := http.ListenAndServeTLS(":808", "go-server.crt", "go-server.key", nil)
	if err != nil {
		log.Fatalf("Server error: %s", err)
	}
}
