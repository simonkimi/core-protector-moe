package net

type Server struct {
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
