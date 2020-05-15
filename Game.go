package main

type Game struct {
	// game uuid which is used mainly for debugging and finding instances of games
	gameUid string

	// Map that the game will be played on.
	gameMap GameMap

	// Starting team 1/2
	startingTeam int

	// The round currently being played
	currentRound int

	// The lockins that will occur for this game.
	team1lockins []Pick
	team2lockins []Pick

	// team ready states
	team1Ready bool
	team2Ready bool

	// team UID's
	team1UID string
	team2UID string

	//GameState, what state the game is currently in
	gameState GameState
}

type GameState string

const (
	created      GameState = "created"
	selectingMap GameState = "selectingMap"
	selectedMap  GameState = "selectedMap"
	readied      GameState = "readied"
	picking      GameState = "picking"
	picked       GameState = "picked"
	finished     GameState = "finished"
)

func (game Game) checkReadyStatus() (bool, error) {

	// Checking the status
	if game.gameState != created {

	}

	return true, nil

}

func (game Game) readyTeam(teamNumber int) {

	// checks the ready status of a team once a team has been readied up
	game.checkReadyStatus()
}
