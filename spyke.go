package main

import (
	"code.google.com/p/go.net/websocket"
	"code.google.com/p/gorest"
	"code.google.com/p/gorilla/mux"
	"code.google.com/p/goweb/goweb"

	restful "github.com/emicklei/go-restful"

	"flag"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

var httpFlag *string = flag.String("http", ":80", "listen [http] hostname:port or ip:port")

func main() {
	flag.Parse()
}

func wsHandler(ws *websocket.Conn) {
	u := &User{}
	uc := &UserConn{
		conn: ws,
	}

	userRegister(u, uc)
	defer userUnregister(u)

	url, _ := url.Parse(ws.LocalAddr().String())
	//roomId, err := strconv.Atoi(url.Query().Get("roomId"))
	_, err := strconv.Atoi(url.Query().Get("roomId"))
	if err != nil {
		return
	}

	/*
		// TODO
		// FIX
		room := getRoom(roomId)
		if room == nil {
			return
		}

		room.addUser(uc)
		defer room.removeUser(uc)

	*/

	// enroll the user conn in the map

	go func() {
		for {
			msg := new(Msg)
			if err := websocket.JSON.Receive(ws, msg); err != nil {
				uc.deadc <- 1
				return
			}
			msg.From = u.Id
			userConnMap[msg.To].out <- msg // this is likely not threadsafe
		}
	}()

	for {
		select {
		case msg := <-uc.out:
			if err := websocket.JSON.Send(ws, msg); err != nil {
				break
			}
		case _ = <-uc.deadc:
			return
		}
	}
}
