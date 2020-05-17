"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
// @ts-ignore
const express_1 = __importDefault(require("express"));
const Lobby_1 = require("./classes/Lobby");
const app = express_1.default();
app.set("port", process.env.PORT || 3000);
let http = require("http").Server(app);
let io = require("socket.io")(http);
let lobbies = [];
io.on("connection", (socket) => {
    console.log("a user connected");
    lobbies.push(new Lobby_1.Lobby(3, 1, 'Cloud9', 'UCAP'));
});
io.on("createLobby", () => {
});
const server = http.listen(8000, () => {
    console.log("listening on *:3000");
});
//# sourceMappingURL=server.js.map