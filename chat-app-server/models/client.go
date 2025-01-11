package models

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

const (
	TypeJoin  = "join"
	TypeLeave = "leave"
	TypeChat  = "chat"
)

type Message struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type ChatPayload struct {
	UserName string `json:"userName"`
	Text     string `json:"text"`
}

type Client struct {
	Conn     *websocket.Conn
	UserName string
	Room     *Room
	Send     chan []byte
}

func NewClient(conn *websocket.Conn, userName string, room *Room) *Client {
	client := &Client{
		Conn:     conn,
		UserName: userName,
		Room:     room,
		Send:     make(chan []byte),
	}

	go client.WriteMessages()

	return client
}

func (c *Client) ReadMessages() {
	defer func() {
		c.Room.Leave <- c
		c.Conn.Close()
	}()

	for {
		_, rawMsg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}

		var message Message
		if err := json.Unmarshal(rawMsg, &message); err != nil {
			log.Printf("Error parsing message: %v", err)
			continue
		}

		switch message.Type {
		case TypeChat:
			payload := ChatPayload{
				UserName: c.UserName,
				Text:     message.Payload.(string),
			}
			message := Message{
				Type:    TypeChat,
				Payload: payload,
			}
			messageJSON, err := json.Marshal(message)
			if err != nil {
				log.Printf("Error serializing chat message: %v", err)
				continue
			}
			c.Room.Broadcast <- messageJSON

		case TypeJoin:
			log.Printf("%s joined room %s", c.UserName, c.Room.Name)

		case TypeLeave:
			log.Printf("%s left room %s", c.UserName, c.Room.Name)
			return

		default:
			log.Printf("Unknown message type: %s", message.Type)
		}
	}
}

func (c *Client) WriteMessages() {
	defer c.Conn.Close()

	for msg := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Printf("Error writing message: %v", err)
			break
		}
	}
}
