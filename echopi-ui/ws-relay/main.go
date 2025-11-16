package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = map[*websocket.Conn]bool{}

func main() {

	//===============================
	// WS: Browser UI connects here
	//===============================
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("❌ WS upgrade error:", err)
			return
		}

		clients[conn] = true
		fmt.Println("🌐 UI connected")

		// Reader loop (ignore empty WS frames from browser)
		go func(c *websocket.Conn) {
			for {
				_, msg, err := c.ReadMessage()
				if err != nil {
					fmt.Println("🔌 UI disconnected")
					delete(clients, c)
					c.Close()
					return
				}

				trim := strings.TrimSpace(string(msg))

				// Drop Chrome's blank frames
				if trim == "" {
					continue
				}
			}
		}(conn)
	})

	//====================================
	// POST /forward → Broadcast to UI(s)
	//====================================
	http.HandleFunc("/forward", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "bad body", 400)
			return
		}

		raw := strings.TrimSpace(string(body))

		// Drop empty JSON or accidental whitespace
		if raw == "" || raw == "{}" {
			w.Write([]byte("ignored"))
			return
		}

		// Broadcast EXACT JSON body to all connected UIs
		for c := range clients {
			err := c.WriteMessage(websocket.TextMessage, []byte(raw))
			if err != nil {
				fmt.Println("❌ WS write error:", err)
				c.Close()
				delete(clients, c)
			}
		}

		w.Write([]byte("ok"))
	})

	fmt.Println("🚀 WebSocket Relay Server active on :9001")
	log.Fatal(http.ListenAndServe("0.0.0.0:9001", nil))
}
