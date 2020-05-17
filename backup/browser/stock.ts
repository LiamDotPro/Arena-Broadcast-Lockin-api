/* Do not change, this code is generated from Golang structs */


export class CreateLobbyPayload {
    MaxRounds: number;
    StartingTeam: number;
    Team1Name: string;
    Team2Name: string;

    static createFrom(source: any) {
        if ('string' === typeof source) source = JSON.parse(source);
        const result = new CreateLobbyPayload();
        result.MaxRounds = source["MaxRounds"];
        result.StartingTeam = source["StartingTeam"];
        result.Team1Name = source["Team1Name"];
        result.Team2Name = source["Team2Name"];
        return result;
    }

}
export class JoinTeamPayload {
    LobbyUID: string;
    TeamUID: string;

    static createFrom(source: any) {
        if ('string' === typeof source) source = JSON.parse(source);
        const result = new JoinTeamPayload();
        result.LobbyUID = source["LobbyUID"];
        result.TeamUID = source["TeamUID"];
        return result;
    }

}
export class Pick {
    Character: string;
    Spec: string;

    static createFrom(source: any) {
        if ('string' === typeof source) source = JSON.parse(source);
        const result = new Pick();
        result.Character = source["Character"];
        result.Spec = source["Spec"];
        return result;
    }

}
export class Game {
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

    static createFrom(source: any) {
        if ('string' === typeof source) source = JSON.parse(source);
        const result = new Game();
        result.GameUid = source["GameUid"];
        result.GameMap = source["GameMap"];
        result.StartingTeam = source["StartingTeam"];
        result.CurrentRound = source["CurrentRound"];
        result.Team1lockins = source["Team1lockins"] ? source["Team1lockins"].map(function(element: any) { return Pick.createFrom(element); }) : null;
        result.Team2lockins = source["Team2lockins"] ? source["Team2lockins"].map(function(element: any) { return Pick.createFrom(element); }) : null;
        result.Team1Ready = source["Team1Ready"];
        result.Team2Ready = source["Team2Ready"];
        result.Team1UID = source["Team1UID"];
        result.Team2UID = source["Team2UID"];
        result.GameState = source["GameState"];
        result.PickingStrategy = source["PickingStrategy"];
        return result;
    }

}
export class Lobby {
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

    static createFrom(source: any) {
        if ('string' === typeof source) source = JSON.parse(source);
        const result = new Lobby();
        result.Team1UID = source["Team1UID"];
        result.Team2UID = source["Team2UID"];
        result.Team1Name = source["Team1Name"];
        result.Team2Name = source["Team2Name"];
        result.Team1link = source["Team1link"];
        result.Team2link = source["Team2link"];
        result.Games = source["Games"] ? source["Games"].map(function(element: any) { return Game.createFrom(element); }) : null;
        result.Enabled = source["Enabled"];
        result.LobbyUID = source["LobbyUID"];
        result.MaxRounds = source["MaxRounds"];
        result.StartingTeam = source["StartingTeam"];
        return result;
    }

}