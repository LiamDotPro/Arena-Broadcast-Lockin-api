import { GameLobby } from '../interfaces/GameLobby'
import { Game } from '../interfaces/Game'
import { UUID } from '../shared/UUID'

export class Lobby implements GameLobby {

    Enabled: boolean = false
    Games: Game[] = []
    LobbyUID: string = ''
    MaxRounds: number = 0
    StartingTeam: number = 0

    Team1Name: string = ''
    Team1UID: string = ''
    Team1link: string = ''

    Team2Name: string = ''
    Team2UID: string = ''
    Team2link: string = ''

    constructor(maxRounds: number, startingTeam: number, team1Name: string, team2Name: string) {

        const lobbyUUID = UUID()
        const team1UID = UUID()
        const team2UID = UUID()

        this.LobbyUID = lobbyUUID

        this.Team1Name = team1Name
        this.Team1UID = team1UID
        this.Team1link = `/pv/${ lobbyUUID }/${ team1UID }`

        this.Team2Name = team2Name
        this.Team2UID = team2UID
        this.Team2link = `/pv/${ lobbyUUID }/${ team2UID }`

        this.StartingTeam = startingTeam
        this.MaxRounds = maxRounds

    }


}
