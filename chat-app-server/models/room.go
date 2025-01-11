package models

import "log"

type Room struct {
	Name      string
	Clients   map[*Client]bool
	Join      chan *Client
	Leave     chan *Client
	Broadcast chan []byte
}

func NewRoom(name string) *Room {
	return &Room{
		Name:      name,
		Clients:   make(map[*Client]bool),
		Join:      make(chan *Client),
		Leave:     make(chan *Client),
		Broadcast: make(chan []byte),
	}
}

func (r *Room) Run() {
	for {
		select {
		case client := <-r.Join:
			r.Clients[client] = true
			log.Printf("%s joined room %s", client.UserName, r.Name)

		case client := <-r.Leave:
			delete(r.Clients, client)
			close(client.Send)
			log.Printf("%s left room %s", client.UserName, r.Name)

		case message := <-r.Broadcast:
			for client := range r.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(r.Clients, client)
				}
			}
		}
	}
}
