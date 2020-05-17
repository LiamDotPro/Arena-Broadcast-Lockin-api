package backup

import "github.com/tkrajina/typescriptify-golang-structs/typescriptify"

func generateInterfaces() {
	converter := typescriptify.New()

	converter.CreateFromMethod = true
	converter.Indent = "    "
	converter.BackupDir = ""

	converter.Add(CreateLobbyPayload{})
	converter.Add(JoinTeamPayload{})
	converter.Add(Lobby{})

	err := converter.ConvertToFile("browser/stock.ts")

	if err != nil {
		panic(err.Error())
	}

	converter.CreateInterface = true

	err = converter.ConvertToFile("browser/interfaces.ts")

	if err != nil {
		panic(err.Error())
	}


}
