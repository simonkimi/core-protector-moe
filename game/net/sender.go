package net

import (
	"core-protector-moe/game/bean"
	"encoding/json"
	"fmt"
)

func (server ServerInfo) GetGameVersion() (bean.LoginVersionBean, error) {
	var loginVersion bean.LoginVersionBean
	netData, err := server.Http.Get(server.UrlVersion)
	if err != nil {
		return loginVersion, err
	}
	err = json.Unmarshal(netData.Body, &loginVersion)
	return loginVersion, nil
}

func (server ServerInfo) LoginLogin(data string) (bean.LoginGetLoginBean, error) {
	var loginLogin bean.LoginGetLoginBean
	url := server.LoginApiHead + "1.0/get/login/@self"
	netData, err := server.Http.Post(url, data)
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
