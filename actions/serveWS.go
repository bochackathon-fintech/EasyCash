package actions

import (
	"log"

	"github.com/gobuffalo/buffalo"
)

// serveWS default implementation.
func serveWS(c buffalo.Context) error {
	log.Println("Serving Websocket for ", c.Request().RemoteAddr)
	//Websocket
	hub := newHub()
	go hub.run()
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println(err)
		return err
	}
	client := &Client{hub: hub, conn: conn, c: c, send: make(chan []byte, 256)}
	client.hub.register <- client
	go client.writePump()
	client.readPump()
	return err
}
