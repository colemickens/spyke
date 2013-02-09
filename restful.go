package main

import (
	restful "github.com/emicklei/go-restful"

	"log"
)

func getAllRooms(req *restful.Request, resp *restful.Response) {

}

func getRoom(req *restful.Request, resp *restful.Response) {

}

func getRoomUsers(req *restful.Request, resp *restful.Response) {

}

func NewRoomService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v2/room").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(
		ws.Get("").
		To(getAllRooms).
		Doc("get all the rooms"))

	ws.Route(
		ws.Get("/{id}").
		To(getRoom).
		Doc("get a room by its id").
		Param(ws.PathParameter("id", "room id"))

	ws.Route(
		ws.Get("/{id}/users").
		To(getRoomUsers).
		Doc("get a room's users by the room id").
		Param(ws.PathParameter("id", "room id"))
}