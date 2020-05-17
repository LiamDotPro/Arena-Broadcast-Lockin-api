package backup

type Game struct {
	// game uuid which is used mainly for debugging and finding instances of games
	GameUid string `json:"GameUid"`

	// Map that the game will be played on.
	GameMap GameMap `json:"GameMap"`

	// Starting team 1/2
	StartingTeam int `json:"StartingTeam"`

	// The round currently being played
	CurrentRound int `json:"CurrentRound"`

	// The lockins that will occur for this game.
	Team1lockins []Pick `json:"Team1lockins"`
	Team2lockins []Pick `json:"Team2lockins"`

	// team ready states
	Team1Ready bool `json:"Team1Ready"`
	Team2Ready bool `json:"Team2Ready"`

	// team UID's
	Team1UID string `json:"Team1UID"`
	Team2UID string `json:"Team2UID"`

	//GameState, what state the game is currently in
	GameState GameState `json:"GameState"`

	// Picking Strategy
	PickingStrategy PickingStrategy `json:"PickingStrategy"`
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

type PickingStrategy string

const (
	blind     PickingStrategy = "blind"
	turnBased PickingStrategy = "turnBased"
)

func (game Game) checkReadyStatus() (bool, error) {

	// Checking the status
	if game.GameState != created {

	}

	return true, nil

}

func (game Game) readyTeam(teamNumber int) {

	// checks the ready status of a team once a team has been readied up
	game.checkReadyStatus()
}
