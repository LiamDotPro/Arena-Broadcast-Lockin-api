package main

import (
	"github.com/k0kubun/pp"
	"testing"
)

func TestCreateLobby(t *testing.T) {

	maxRounds := 3
	startingTeam := 1
	team1Name := "Cloud9"
	team2Name := "UCAP eSports"

	lobby := createLobby(maxRounds, startingTeam, team1Name, team2Name)

	if lobby.startingTeam != startingTeam {
		t.Errorf("starting team was not correctly set..: %v", startingTeam)
	}

	if lobby.maxRounds != maxRounds {
		t.Errorf("max rounds was not correctly set..: %v", maxRounds)
	}

	pp.Print(lobby)

}
