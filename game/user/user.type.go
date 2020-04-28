package user

import "liverProtectorMoe/game/net"

type Base struct {
	User struct {
		username   string
		password   string
		ServerType int
	}
	Server net.Server
}
