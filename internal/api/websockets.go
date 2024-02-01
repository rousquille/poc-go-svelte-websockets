package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"time"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type WsManager struct {
	clients WsClientList
	sync.RWMutex
}

type WsClientList map[*WsClient]bool

type WsClient struct {
	connection *websocket.Conn
	wsManager  *WsManager
	writeChan  chan []byte
}

type WsMessage struct {
	Type    string `json:"type"`
	Payload any    `json:"payload"`
}

func NewWsManager() *WsManager {
	return &WsManager{
		clients: make(WsClientList),
	}
}

func NewWsClient(conn *websocket.Conn, wsManager *WsManager) *WsClient {
	return &WsClient{
		connection: conn,
		wsManager:  wsManager,
		writeChan:  make(chan []byte),
	}
}

func (w *WsManager) addWsClient(client *WsClient) {
	w.Lock()
	defer w.Unlock()

	w.clients[client] = true
}

func (w *WsManager) removeWsClient(client *WsClient) {
	w.Lock()
	defer w.Unlock()

	_, ok := w.clients[client]
	if ok {
		client.connection.Close()
		delete(w.clients, client)
	}
}

func (w *WsClient) readMessages() {
	defer func() {
		w.wsManager.removeWsClient(w)
	}()

	for {
		messageType, payload, err := w.connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message : %v\n", err)
			}
			break
		}
		log.Println("MessageType :", messageType)
		log.Println("Payload :", string(payload))

		var wsMessage WsMessage

		err = json.Unmarshal(payload, &wsMessage)
		if err != nil {
			log.Printf("error to decode message : %v\n", err)
			break
		}

		switch {
		case wsMessage.Type == "chat":
			payloadUpdated, err := addTimestampChatMsgAndEncode(wsMessage)
			if err != nil {
				log.Printf("error to re-encode message : %v\n", err)
				break
			}
			w.broadcastingMessage(payloadUpdated)

		default:
			fmt.Println("unrecognized message type")
			break
		}
	}
}

func (w *WsClient) writeMessages() {
	defer func() {
		w.wsManager.removeWsClient(w)
	}()

	for {
		select {
		case message, ok := <-w.writeChan:

			if !ok {
				err := w.connection.WriteMessage(websocket.CloseMessage, nil)
				if err != nil {
					log.Println("connection closed : ", err)
				}
				return
			}

			err := w.connection.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println(err)
			}
			log.Println("sent message :", string(message))
		}
	}
}

func (w *WsClient) broadcastingMessage(payload []byte) {
	for wsClient := range w.wsManager.clients {
		wsClient.writeChan <- payload
	}
}

func addTimestampChatMsgAndEncode(wsMessage WsMessage) ([]byte, error) {
	payloadChat, ok := wsMessage.Payload.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("type assertion error for payloadChat %v\n", payloadChat)
	}

	payloadChat["timestamp"] = time.Now().Unix()

	encMsg, err := json.Marshal(WsMessage{Type: "chat", Payload: payloadChat})
	if err != nil {
		return nil, err
	}
	return encMsg, nil
}
