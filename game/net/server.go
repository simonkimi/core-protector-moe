package net

import (
	"core-protector-moe/game/user"
	"fmt"
)

/**
登录游戏
*/
func Login(base *user.Base) error {
	// 取游戏版本号, 初始化URL数据
	if err := initVersion(base); err != nil {
		fmt.Println("初始化服务器失败")
		return err
	}
	// 获取游戏的Token值
	if err := initToken(base); err != nil {
		fmt.Println("获取游戏Token失败")
		return err
	}
	return nil
}

func initVersion(base *user.Base) error {
	// 取游戏版本号
	if base.Server.FirstLogin {
		return nil
	}
	gameVersionData, err := GetGameVersion(base)
	if err != nil {
		return err
	}
	base.Server.UrlVersion = gameVersionData.Version.NewVersionId
	base.Server.ResVersion = gameVersionData.Version.DataVersion
	base.Server.LoginHead = gameVersionData.LoginServer
	base.Server.LoginApiHead = gameVersionData.HmLoginServer
	base.Server.FirstLogin = true
	return nil
}

func initToken(base *user.Base) error {
	if base.Server.Token != "" {
		return nil
	}
	getLoginToken(base)
	return nil
}

func getLoginToken(base *user.Base) (string, error) {
	loginJson := fmt.Sprintf("{"+
		"\"platform\": \"0\","+
		"\"appid\": \"0\","+
		"\"app_server_type\": \"0\","+
		"\"password\": \"%s\","+
		"\"username\": \"%s\""+
		"}", base.User.Password, base.User.Username)
	data, err := LoginLogin(base, loginJson)
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
