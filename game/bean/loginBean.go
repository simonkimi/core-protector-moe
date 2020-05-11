package bean

type Version struct {
	NewVersionId string `json:"newVersionId"`
	DataVersion  string
}

type LoginVersionBean struct {
	Eid           string `json:"eid"`
	ResUrlWu      string
	ResUrl        string
	ResVersion    string
	Version       Version `json:"version"`
	LoginServer   string  `json:"loginServer"`
	HmLoginServer string  `json:"hmLoginServer"`
	DataVersion   string
}

type LoginBean struct {
	Error        string `json:"error"`
	Errmsg       string `json:"errmsg"`
	Access_token string `json:"access_token"`
	Token        string `json:"token"`
}

type ServerList struct {
	Id   string `json:"id"`
	Host string `json:"host"`
	Name string `json:"name"`
}

type LoginServerListBean struct {
	UserId        string       `json:"userId"`
	DefaultServer string       `json:"defaultServer"`
	ServerList    []ServerList `json:"serverList"`
}

type LoginGetLoginBean struct {
	Access_token string `json:"access_token"`
	Error        int    `json:"error"`
}
