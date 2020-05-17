package backup

type JoinTeamPayload struct {
	LobbyUID string `json:"LobbyUID"`
	TeamUID  string `json:"TeamUID"`
}

type CreateLobbyPayload struct {
	MaxRounds    int    `json:"MaxRounds"`
	StartingTeam int    `json:"StartingTeam"`
	Team1Name    string `json:"Team1Name"`
	Team2Name    string `json:"Team2Name"`
}
