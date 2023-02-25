package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (c *Config)routes() http.Handler{
	chiMux := chi.NewRouter()

	// specify who is allowed to connect
	chiMux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF_Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// check if service is still alive
	chiMux.Use(middleware.Heartbeat("/ping"))

	// routes
	chiMux.Get("/setup-ws", c.setUpWebSocketConn)
	chiMux.Get("/create-room", c.createRoom)
	chiMux.Get("/join-room", c.joinRoom)

	return chiMux
}