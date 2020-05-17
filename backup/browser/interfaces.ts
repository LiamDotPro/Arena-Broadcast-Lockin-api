/* Do not change, this code is generated from Golang structs */


export interface CreateLobbyPayload {
    MaxRounds: number;
    StartingTeam: number;
    Team1Name: string;
    Team2Name: string;
}
export interface JoinTeamPayload {
    LobbyUID: string;
    TeamUID: string;
}
export interface Pick {
    Character: string;
    Spec: string;
}
export interface Game {
    GameUid: string;
    GameMap: string;
    StartingTeam: number;
    CurrentRound: number;
    Team1lockins: Pick[];
    Team2lockins: Pick[];
    Team1Ready: boolean;
    Team2Ready: boolean;
    Team1UID: string;
    Team2UID: string;
    GameState: string;
    PickingStrategy: string;
}
export interface Lobby {
    Team1UID: string;
    Team2UID: string;
    Team1Name: string;
    Team2Name: string;
    Team1link: string;
    Team2link: string;
    Games: Game[];
    Enabled: boolean;
    LobbyUID: string;
    MaxRounds: number;
    StartingTeam: number;
}