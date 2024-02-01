package api

import (
	"log"
	"net/http"
)

func (s *Server) wsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wsUpgrader.CheckOrigin = func(r *http.Request) bool { return true }

		wsConn, err := wsUpgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Println("client connected")

		client := NewWsClient(wsConn, s.wsManager)
		s.wsManager.addWsClient(client)

		go client.readMessages()
		go client.writeMessages()
	}
}
