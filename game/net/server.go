package net

import (
	"fmt"
)

/**
登录游戏
*/
func (server *ServerInfo) Login() error {
	// 取游戏版本号, 初始化URL数据
	if err := server.initVersion(); err != nil {
		fmt.Println("初始化服务器失败")
		return err
	}
	// 获取游戏的Token值
	if err := server.initToken(); err != nil {
		fmt.Println("获取游戏Token失败")
		return err
	}
	return nil
}

func (server *ServerInfo) initVersion() error {
	// 取游戏版本号
	if server.FirstLogin {
		return nil
	}
	gameVersionData, err := server.GetGameVersion()
	if err != nil {
		return err
	}
	server.UrlVersion = gameVersionData.Version.NewVersionId
	server.ResVersion = gameVersionData.Version.DataVersion
	server.LoginHead = gameVersionData.LoginServer
	server.LoginApiHead = gameVersionData.HmLoginServer
	server.FirstLogin = true
	return nil
}

func (server *ServerInfo) initToken() error {
	if server.Token != "" {
		return nil
	}
	server.getLoginToken()
	return nil
}

func (server *ServerInfo) getLoginToken() (string, error) {
	loginJson := fmt.Sprintf("{"+
		"\"platform\": \"0\","+
		"\"appid\": \"0\","+
		"\"app_server_type\": \"0\","+
		"\"password\": \"%s\","+
		"\"username\": \"%s\""+
		"}", server.User.Password, server.User.Username)
	data, err := server.LoginLogin(loginJson)
	if err != nil {
		return "", err
	}
	if data.Error != 0 {
		err := &LoginError{code: data.Error, message: "获取Token出错"}
		fmt.Println("getLoginToken", err)
		return "", err
	}
	return data.Access_token, nil
}
