package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/ryanjoy0000/chat-room/chat-service/errPkg"
	"github.com/google/uuid"
	"github.com/ryanjoy0000/chat-room/chat-service/models"
)

var origins = []string{ "http://127.0.0.1", "http://localhost"}


var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,

	 // Resolve cross-domain problems
	CheckOrigin: func(r *http.Request) bool {
        var origin = r.Header.Get("origin")
        for _, allowOrigin := range origins {
            if origin == allowOrigin {
                return true
            }
        }
        return false
    },
}


func (c *Config)setUpWebSocketConn(w http.ResponseWriter, r *http.Request){
	log.Println("setting up websocket connection")

	// Upgrade the HTTP server connection to the WebSocket protocol.
	conn, err := upgrader.Upgrade(w, r, nil)
	errPkg.HandleHttpErr(err, "ws conn err", w, http.StatusInternalServerError)

	log.Println("websocket connection established!")

	// defers & closes the underlying network connection
	defer conn.Close()

	// // repeat
	// for {
	// 	// returns the next data message received from the peer		
	// 	msgType, msgBSlice, err := conn.ReadMessage()
	// 	errPkg.HandleHttpErr(err, "ws conn read err", w, http.StatusInternalServerError)
	// 	if err != nil {break}
	// 	log.Println("received: ",msgBSlice)
	// 	// returns a writer for the next message to send.
	// 	err = conn.WriteMessage(msgType, msgBSlice)
	// 	errPkg.HandleHttpErr(err, "ws conn write err", w, http.StatusInternalServerError)
	// 	if err != nil {break}
	// 	log.Println("sent msg")
	// }


}


func (c *Config)createRoom(w http.ResponseWriter, r *http.Request){
	log.Println("before creating room, roomlist:", app.RoomList)
	
	// create room id
	roomId := uuid.New().String()
	
	// create member id
	memberId := uuid.New().String()

	// create member as a host
	m1 := models.Member{
		Host: true,
		Id: memberId,
	}

	// create member list
	mList := []models.Member{m1}

	// add member list to room map, with room id
	app.RoomList[roomId] = mList

	log.Println("after creating room, roomlist:", app.RoomList)
	
	// create room join link
	msg := "Room created. Room Link : " + "{your domain}/join-room?id=" + roomId
	fmt.Fprintf(w, msg)

}

func (c *Config)joinRoom(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "joining room")
}