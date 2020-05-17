package backup

import (
	"encoding/json"
	"fmt"
	"github.com/googollee/go-socket.io"
	"github.com/k0kubun/pp"
	"log"
	"net/http"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, PATCH, GET, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", allowHeaders)

		next.ServeHTTP(w, r)
	})
}


// err = json.Unmarshal([]byte(msg), &socketData)
func main() {

	server, err := socketio.NewServer(nil)

	if err != nil {
		log.Fatal(err)
	}

	//-------
	// Player view

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	// Joins someone to a player view
	server.OnEvent("/", "joinTeam", func(s socketio.Conn, payload string) error {
		pp.Println("A player is attempting to join a team")

		joinTeamPayload := JoinTeamPayload{}

		err = json.Unmarshal([]byte(payload), &joinTeamPayload)

		pp.Println(joinTeamPayload)

		err, lobby := getLobbyByUID(joinTeamPayload.LobbyUID)

		// Validate that the lobby exists
		if err != nil {
			pp.Println(err)
			return nil
		}

		// Assign socket for future interactions
		outcome := lobby.assignSocketToTeam(s, joinTeamPayload.TeamUID)

		if outcome != nil {
			pp.Println(outcome)
			return nil
		}

		pp.Println("Lobby added a new connection to a slot")

		// Set the join team payload as the interface for finding information
		s.SetContext(joinTeamPayload)

		pp.Println("Connection has lobby information added")

		s.Emit("successfulJoined", "")

		return nil
	})

	server.OnEvent("/", "requestCurrentGameData", func(s socketio.Conn, payload string) error {

		pp.Println("A player is requesting the current state for UI setup")

		joinTeamPayload := JoinTeamPayload{}

		err = json.Unmarshal([]byte(payload), &joinTeamPayload)

		pp.Println(joinTeamPayload)

		//  Get the current lobby in question
		err, lobby := getLobbyByUID(joinTeamPayload.LobbyUID)

		// Validate that the lobby exists
		if err != nil {
			pp.Println(err)
			return nil
		}

		value, _ := json.Marshal(&lobby.Games)

		s.Emit("receiveGameData", string(value))

		return nil
	})

	//------------

	server.OnError("/", func(s socketio.Conn, e error) {
		// fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		pp.Println("closed", reason)
		pp.Println(s.Context())
	})

	//-----------

	// Joins someone to the admin room
	server.OnEvent("/", "joinAdmin", func(s socketio.Conn) error {
		// Need to add some validation here.

		// Join them to the control room
		s.Join("control")

		value, _ := json.Marshal(&Lobbies)

		s.Emit("adminSetup", string(value))

		return nil
	})

	// Admin Events
	server.OnEvent("/", "createLobby", func(s socketio.Conn, payload string) error {

		// A request for lobby creation has been found.
		createLobbyPayload := CreateLobbyPayload{}

		err = json.Unmarshal([]byte(payload), &createLobbyPayload)

		newLobby := createLobby(
			createLobbyPayload.MaxRounds,
			createLobbyPayload.StartingTeam,
			createLobbyPayload.Team1Name,
			createLobbyPayload.Team2Name,
		)

		// Append the lobby directly into the global lobby state..
		Lobbies = append(Lobbies, newLobby)

		value, _ := json.Marshal(&newLobby)

		server.BroadcastToRoom("/", "control", "newLobbyCreated", string(value))

		return nil
	})


	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", corsMiddleware(server))

	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
