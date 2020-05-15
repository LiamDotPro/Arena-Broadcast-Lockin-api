package main

import (
	"fmt"
	"github.com/lithammer/shortuuid"
)

var Lobbies []Lobby

type Lobby struct {
	// UIDs - used during pv to represent the area they can control.
	team1UID string
	team2UID string

	// The displayed names on the widget
	team1Name string
	team2Name string

	// Storage and concerns for links to given to players.
	team1link string
	team2link string

	// Games
	games []Game

	// Enabled - acts as a flag to override input from players and should be set by an admin
	enabled bool

	// lobbyUid
	lobbyUID string

	// Amount of rounds per team bo1,bo3, bo5
	maxRounds int

	// Which team picks first
	startingTeam int
}

type createLobbyPayload struct {
	MaxRounds    int    `json:"MaxRounds"`
	StartingTeam int    `json:"StartingTeam"`
	Team1Name    string `json:"Team1Name"`
	Team2Name    string `json:"Team2Name"`
}

/**
 * Creates a new Lobby..
 */
func createLobby(maxRounds int, startingTeam int, team1Name string, team2Name string) Lobby {

	// Firstly generate uuids
	lobbyUUID := generateUUID()
	team1UIDGenerated := generateUUID()
	team2UIDGenerated := generateUUID()

	lobby := Lobby{
		lobbyUID:     lobbyUUID,
		team1UID:     team1UIDGenerated,
		team1Name:    team1Name,
		team2Name:    team2Name,
		team2UID:     team2UIDGenerated,
		team1link:    fmt.Sprintf(`/pv/%v/%v`, lobbyUUID, team1UIDGenerated),
		team2link:    fmt.Sprintf(`/pv/%v/%v`, lobbyUUID, team2UIDGenerated),
		maxRounds:    maxRounds,
		enabled:      false,
		startingTeam: startingTeam,
	}

	// By default whenever a new lobby is generated then a game on nagrand is added.
	// With the exception of there being a potential different map for single play.
	if maxRounds != 1 {
		lobby.createDefaultGame(startingTeam)
	}

	if maxRounds == 1 {
		lobby.createGame(startingTeam, 1)
	}

	return lobby

}

/**
 * Create's a new game
 */
func (lob *Lobby) createGame(startingTeam int, round int) {

	newGame := Game{
		gameUid:      shortuuid.New(),
		startingTeam: startingTeam,
		currentRound: round,
		team1Ready:   false,
		team2Ready:   false,
		gameState:    created,
		team1UID:     lob.team1UID,
		team2UID:     lob.team2UID,
	}

	lob.games = append(lob.games, newGame)

}

/**
 * Create's a new game with the default ruleset loaded
 */
func (lob *Lobby) createDefaultGame(startingTeam int) {

	newGame := Game{
		gameUid:      shortuuid.New(),
		startingTeam: startingTeam,
		currentRound: 1,
		team1Ready:   false,
		team2Ready:   false,
		gameState:    selectedMap,
		team1UID:     lob.team1UID,
		team2UID:     lob.team2UID,
		gameMap:      nagrand,
	}

	lob.games = append(lob.games, newGame)

}
