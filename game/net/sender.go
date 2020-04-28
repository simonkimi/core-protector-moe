package net

import (
	"encoding/json"
	"liverProtectorMoe/game/bean"
	"liverProtectorMoe/util"
)

func GetGameVersion(urlVersion string) (bean.LoginVersionBean, error) {
	var loginVersion bean.LoginVersionBean
	netString, err := util.Get(urlVersion, nil, commonHeaders, nil)
	if err != nil {
		return loginVersion, err
	}
	err = json.Unmarshal(netString.Body, &loginVersion)
	return loginVersion, nil
}
