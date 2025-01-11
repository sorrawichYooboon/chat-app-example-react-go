package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/sorrawichYooboon/chat-app-server/handlers"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	e.GET("/ws", handlers.WebSocketHandler)

	log.Println("Starting server on :3000")
	err := e.Start(":3000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
