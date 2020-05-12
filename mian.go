package main

import (
	"core-protector-moe/game/user"
	"os"
)

func main() {
	username := os.Getenv("GO_USERNAME")
	password := os.Getenv("GO_PASSWORD")

	userBase := &user.Base{}
	userBase.InitUser(username, password, 1)
	userBase.ServerInfo.Login()
}
