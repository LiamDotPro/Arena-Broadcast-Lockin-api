package main

type Lobby struct {
	// UIDs - used during pv to represent the area they can control.
	team1UID string
	team2UID string

	// lockins
	team1lockins []Pick
	team2lockins []Pick

	// lobbyUid
	lobbyUID string

	// Amount of rounds per team
	MaxRounds int

	// Which team picks first
	startingTeam int
}

type Pick struct {
	character string
	spec      string
}

func createLobby() {
	
}
