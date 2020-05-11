package user

import "core-protector-moe/util/httpUtil"

type Sender struct {
	Util *httpUtil.Context
}

type Server struct {
	FirstLogin   bool
	Token        string
	AuthKey      string
	AuthHead     string
	ResUrl       string
	Channel      string
	Version      string
	UrlVersion   string
	ResVersion   string
	Host         string
	LoginHead    string
	LoginApiHead string
	Cookie       map[string]string
}

type Base struct {
	User struct {
		Username   string
		Password   string
		ServerType int
	}
	Server Server
	Sender *Sender
}
