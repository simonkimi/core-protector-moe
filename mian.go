package main

import (
	"liverProtectorMoe/game/net"
	"liverProtectorMoe/game/user"
	"os"
)

func main() {
	username := os.Getenv("GO_USERNAME")
	password := os.Getenv("GO_PASSWORD")
	base := user.InitUser(username, password, 1)
	_ = net.FirstLogin(base.Server)
}
