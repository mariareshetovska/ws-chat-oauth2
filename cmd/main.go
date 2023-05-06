package main

import (
	"log"
	"net/http"
	"ws/config"
	"ws/controller"
	ws "ws/ws"
)

func main() {
	config.LoadEnv()
	s := ws.NewServer()

	http.HandleFunc("/", HandleMain)
	http.HandleFunc("/auth/google/login", controller.HandleGoogleLogin)
	http.HandleFunc("/auth/google/callback", controller.HandleGoogleCallback)

	http.HandleFunc("/auth/linkedin/login", controller.HandleLinkedinLogin)
	http.HandleFunc("/auth/linkedin/callback", controller.HandleLinkedinCallback)

	http.HandleFunc("/echo", s.EchoHandler)
	http.HandleFunc("/chat", s.StaticHandler)

	log.Println("Server started on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server error: %s", err)
	}
}

func HandleMain(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving home page to user %s", r.RemoteAddr)
	http.ServeFile(w, r, "home.html")
}
