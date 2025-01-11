package models

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type Message struct {
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
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}

		message := Message{
			UserName: c.UserName,
			Text:     string(msg),
		}

		messageJSON, err := json.Marshal(message)
		if err != nil {
			log.Printf("Error serializing message: %v", err)
			continue
		}

		c.Room.Broadcast <- messageJSON
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
