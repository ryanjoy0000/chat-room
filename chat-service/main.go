package main

import (
	"log"
	"net/http"

	"github.com/ryanjoy0000/chat-room/chat-service/errPkg"
	"github.com/ryanjoy0000/chat-room/chat-service/models"
)

const webPort = ":8080"

type Config struct {
	RoomList map[string][]models.Member
}

// define app config
var app Config

func main() {
	// set logs
	log.SetFlags(log.Lshortfile)

	// initialise app config
	app.init()

	// set up http server
	log.Println("starting chat-service at port: ", webPort)
	srv := &http.Server{
		Addr:    webPort,
		Handler: app.routes(),
	}

	// start server
	err := srv.ListenAndServe()
	errPkg.HandleErr(err, "err while starting server:")
}

func (c *Config) init() {
	c.RoomList = make(map[string][]models.Member)
}
