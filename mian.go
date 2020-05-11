package main

import (
	"core-protector-moe/game/net"
	"core-protector-moe/game/user"
	"os"
)

func main() {
	username := os.Getenv("GO_USERNAME")
	password := os.Getenv("GO_PASSWORD")
	base := user.InitUser(username, password, 1)
	_ = net.Login(&base)
}
