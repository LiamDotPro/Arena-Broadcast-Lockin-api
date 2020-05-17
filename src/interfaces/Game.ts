import { Pick } from './Pick'
import { GameMaps } from '../enums/GameMaps'

export interface Game {
    GameUid: string;
    GameMap: GameMaps;
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
