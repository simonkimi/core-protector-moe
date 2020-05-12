package user

import "core-protector-moe/util/httpUtil"

func (base *Base) InitUser(username string, password string, serverType int) {
	// 初始化登录数据
	base.ServerInfo.FirstLogin = false
	base.ServerInfo.User.Username = username
	base.ServerInfo.User.Password = password
	base.ServerInfo.Token = ""
	switch serverType {
	case 1:
		base.ServerInfo.AuthHead = "HMS 881d3SlFucX5R5hE"
		base.ServerInfo.AuthKey = "kHPmWZ4zQBYP24ubmJ5wA4oz0d8EgIFe"
		base.ServerInfo.Channel = "100016"
		base.ServerInfo.ResUrl = "http://login.jr.moefantasy.com/index/getInitConfigs/"
		base.ServerInfo.UrlVersion = "http://version.jr.moefantasy.com/index/checkVer/4.1.0/100016/2&version=4.1.0&channel=100016&market=2"
		break
	case 2:
		base.ServerInfo.AuthHead = "HMS 881d3SlFucX5R5hE"
		base.ServerInfo.AuthKey = "kHPmWZ4zQBYP24ubmJ5wA4oz0d8EgIFe"
		base.ServerInfo.Channel = "100015"
		base.ServerInfo.ResUrl = "http://loginios.jr.moefantasy.com/index/getInitConfigs/"
		base.ServerInfo.UrlVersion = "http://version.jr.moefantasy.com/index/checkVer/4.1.0/100015/2&version=4.1.0&channel=100015&market=2"
		break
	}
	// 创建服务对象
	base.ServerInfo.Http = httpUtil.MakeUtil(base.ServerInfo.AuthHead, base.ServerInfo.AuthKey)
}
