"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Lobby = void 0;
const UUID_1 = require("../shared/UUID");
class Lobby {
    constructor(maxRounds, startingTeam, team1Name, team2Name) {
        this.Enabled = false;
        this.Games = [];
        this.LobbyUID = '';
        this.MaxRounds = 0;
        this.StartingTeam = 0;
        this.Team1Name = '';
        this.Team1UID = '';
        this.Team1link = '';
        this.Team2Name = '';
        this.Team2UID = '';
        this.Team2link = '';
        const lobbyUUID = UUID_1.UUID();
        const team1UID = UUID_1.UUID();
        const team2UID = UUID_1.UUID();
        this.LobbyUID = lobbyUUID;
        this.Team1Name = team1Name;
        this.Team1UID = team1UID;
        this.Team1link = `/pv/${lobbyUUID}/${team1UID}`;
        this.Team2Name = team2Name;
        this.Team2UID = team2UID;
        this.Team2link = `/pv/${lobbyUUID}/${team2UID}`;
        this.StartingTeam = startingTeam;
        this.MaxRounds = maxRounds;
    }
}
exports.Lobby = Lobby;
//# sourceMappingURL=Lobby.js.map