import { Game } from './Game'

export interface GameLobby {
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
