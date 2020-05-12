package net

import "core-protector-moe/util/httpUtil"

type ServerInfo struct {
	Http *httpUtil.HttpContext
	User struct {
		Username   string
		Password   string
		ServerType int
	}
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
}
