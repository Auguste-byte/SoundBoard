package ws

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan []byte)
var mu sync.Mutex

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // ⚠️ À restreindre en prod
	},
}

// Gère les connexions entrantes WebSocket
func HandleConnections(w http.ResponseWriter, r *http.Request) {
	log.Println("📡 Tentative de connexion WebSocket...")

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("❌ Échec WebSocket upgrade :", err)
		http.Error(w, "WebSocket upgrade failed", http.StatusInternalServerError)
		return
	}
	defer ws.Close()

	log.Println("✅ Connexion WebSocket établie")

	mu.Lock()
	clients[ws] = true
	mu.Unlock()

	// Gère la déconnexion du client
	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			log.Println("👋 Déconnexion client WebSocket :", err)
			mu.Lock()
			delete(clients, ws)
			mu.Unlock()
			break
		}
	}
}

// Diffuse les messages à tous les clients connectés
func StartBroadcast() {
	for {
		message := <-broadcast
		mu.Lock()
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println("❌ Erreur envoi message :", err)
				client.Close()
				delete(clients, client)
			}
		}
		mu.Unlock()
	}
}

// Permet à ton backend d’envoyer un message à tous les clients
func Send(message []byte) {
	broadcast <- message
}
