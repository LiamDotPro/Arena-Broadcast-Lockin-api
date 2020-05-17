package backup

import (
	"fmt"
	"github.com/googollee/go-socket.io"
	"github.com/lithammer/shortuuid"
	"github.com/pkg/errors"
)

var Lobbies []Lobby

type Lobby struct {
	// UIDs - used during pv to represent the area they can control.
	Team1UID string `json:"Team1UID"`
	Team2UID string `json:"Team2UID"`

	// Connected socket for use when commuting events.
	Team1Socket socketio.Conn
	Team2Socket socketio.Conn

	// The displayed names on the widget
	Team1Name string `json:"Team1Name"`
	Team2Name string `json:"Team2Name"`

	// Storage and concerns for links to given to players.
	Team1link string `json:"Team1link"`
	Team2link string `json:"Team2link"`

	// Games
	Games []Game `json:"Games"`

	// Enabled - acts as a flag to override input from players and should be set by an admin
	Enabled bool `json:"Enabled"`

	// lobbyUid
	LobbyUID string `json:"LobbyUID"`

	// Amount of rounds per team bo1,bo3, bo5
	MaxRounds int `json:"MaxRounds"`

	// Which team picks first
	StartingTeam int `json:"StartingTeam"`
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
		LobbyUID:     lobbyUUID,
		Team1UID:     team1UIDGenerated,
		Team1Name:    team1Name,
		Team2Name:    team2Name,
		Team2UID:     team2UIDGenerated,
		Team1link:    fmt.Sprintf(`/pv/%v/%v`, lobbyUUID, team1UIDGenerated),
		Team2link:    fmt.Sprintf(`/pv/%v/%v`, lobbyUUID, team2UIDGenerated),
		MaxRounds:    maxRounds,
		Enabled:      false,
		StartingTeam: startingTeam,
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
		GameUid:      shortuuid.New(),
		StartingTeam: startingTeam,
		CurrentRound: round,
		Team1Ready:   false,
		Team2Ready:   false,
		GameState:    created,
		Team1UID:     lob.Team1UID,
		Team2UID:     lob.Team2UID,
	}

	lob.Games = append(lob.Games, newGame)

}

/**
 * Create's a new game with the default ruleset loaded
 */
func (lob *Lobby) createDefaultGame(startingTeam int) {

	newGame := Game{
		GameUid:      shortuuid.New(),
		StartingTeam: startingTeam,
		CurrentRound: 1,
		Team1Ready:   false,
		Team2Ready:   false,
		GameState:    selectedMap,
		Team1UID:     lob.Team1UID,
		Team2UID:     lob.Team2UID,
		GameMap:      nagrand,
	}

	lob.Games = append(lob.Games, newGame)

}

func (lob *Lobby) assignSocketToTeam(s socketio.Conn, TeamUID string) error {

	if TeamUID == lob.Team1UID {

		if lob.Team1Socket != nil {
			return errors.New("Someone tried connecting ahead of the first connection for team 1")
		}

		lob.Team1Socket = s

		return nil
	}

	if TeamUID == lob.Team2UID {

		if lob.Team2Socket != nil {
			return errors.New("Someone tried connecting ahead of the first connection for team 2")
		}

		lob.Team2Socket = s

		return nil
	}

	return errors.New("Someone tried connecting who didn't have a correct TeamUID")

}

func getLobbyByUID(lobbyUID string) (error, *Lobby) {

	for i := range Lobbies {
		if Lobbies[i].LobbyUID == lobbyUID {
			return nil, &Lobbies[i]
		}
	}

	return errors.New("A requested lobby could not be found.."), &Lobby{}
}
