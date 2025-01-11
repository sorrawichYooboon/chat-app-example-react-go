package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/sorrawichYooboon/chat-app-server/models"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var rooms = make(map[string]*models.Room)

func WebSocketHandler(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return err
	}

	defer conn.Close()

	userName := c.QueryParam("userName")
	roomName := c.QueryParam("roomName")
	if userName == "" || roomName == "" {
		log.Println("Missing userName or roomName")
		return nil
	}

	room, exists := rooms[roomName]
	if !exists {
		room = models.NewRoom(roomName)
		rooms[roomName] = room
		go room.Run()
	}

	client := models.NewClient(conn, userName, room)
	room.Join <- client

	client.ReadMessages()

	return nil
}
