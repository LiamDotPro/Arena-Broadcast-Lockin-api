package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

func GinMiddleware(allowOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Request.Header.Del("Origin")

		c.Next()
	}
}

type socketData struct {
	Text string
}

// err = json.Unmarshal([]byte(msg), &socketData)
func main() {

	server, err := socketio.NewServer(nil)

	router := gin.New()

	if err != nil {
		log.Fatal(err)
	}

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	// Joins someone to a player view
	server.OnEvent("/", "join", func(s socketio.Conn, key string) error {
		fmt.Println("Reached here..")
		return nil
	})

	// Joins someone to the admin room
	server.OnEvent("/", "joinAdmin", func(s socketio.Conn) error {
		fmt.Println("An admin has connected to the control panel..")
		s.Join("admin")
		server.BroadcastToRoom("/", "admin", "adminSetup", Lobbies)
		return nil
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		// fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	// Admin Events
	server.OnEvent("/admin", "createLobby", func(s socketio.Conn, payload string) error {

		// A request for lobby creation has been found.
		createLobbyPayload := createLobbyPayload{}

		err = json.Unmarshal([]byte(payload), &createLobbyPayload)

		// Append the lobby directly into the global lobby state..
		Lobbies = append(Lobbies, createLobby(
			createLobbyPayload.MaxRounds,
			createLobbyPayload.StartingTeam,
			createLobbyPayload.Team1Name,
			createLobbyPayload.Team2Name,
		))

		server.BroadcastToRoom("/", "control", "newLobbyCreated", "test")

		return nil
	})


	go server.Serve()
	defer server.Close()

	router.Use(GinMiddleware("http://localhost:3000"))
	router.GET("/socket.io/*any", gin.WrapH(server))
	router.POST("/socket.io/*any", gin.WrapH(server))
	router.StaticFS("/public", http.Dir("../asset"))

	_ = router.Run(":8000")
}
