package net

import (
	"core-protector-moe/game/bean"
	"core-protector-moe/game/user"
	"encoding/json"
	"fmt"
)

func GetGameVersion(base *user.Base) (bean.LoginVersionBean, error) {
	var loginVersion bean.LoginVersionBean
	netData, err := base.Sender.Util.Get(base.Server.UrlVersion)
	if err != nil {
		return loginVersion, err
	}
	err = json.Unmarshal(netData.Body, &loginVersion)
	return loginVersion, nil
}

func LoginLogin(base *user.Base, data string) (bean.LoginGetLoginBean, error) {
	var loginLogin bean.LoginGetLoginBean
	url := base.Server.LoginApiHead + "1.0/get/login/@self"
	netData, err := base.Sender.Util.Post(url, data)
	if err != nil {
		return loginLogin, err
	}
	err = json.Unmarshal(netData.Body, &loginLogin)
	if err != nil {
		fmt.Println("解析Json失败", err)
		return loginLogin, err
	}
	return loginLogin, nil
}
