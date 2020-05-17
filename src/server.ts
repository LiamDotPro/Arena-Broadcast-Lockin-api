// @ts-ignore
import express from "express";
import { GameLobby } from './interfaces/GameLobby'
import { Lobby } from './classes/Lobby'

const app = express();
app.set("port", process.env.PORT || 3000);

let http = require("http").Server(app);
let io = require("socket.io")(http);

let lobbies: GameLobby[] = []

io.on("connection", (socket: any) => {
    console.log("a user connected");
    lobbies.push(new Lobby(3, 1, 'Cloud9', 'UCAP'))
})

io.on("createLobby", () => {

})



const server = http.listen(8000, () => {
    console.log("listening on *:3000");
})
