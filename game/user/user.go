package user

func InitUser(username string, password string, serverType int) Base {
	var base Base
	base.User.username = username
	base.User.password = password
	switch serverType {
	case 1:
		base.Server.AuthHead = "HMS 881d3SlFucX5R5hE"
		base.Server.AuthKey = "kHPmWZ4zQBYP24ubmJ5wA4oz0d8EgIFe"
		base.Server.Channel = "100016"
		base.Server.ResUrl = "http://login.jr.moefantasy.com/index/getInitConfigs/"
		base.Server.UrlVersion = "http://version.jr.moefantasy.com/index/checkVer/4.1.0/100016/2&version=4.1.0&channel=100016&market=2"
		break
	case 2:
		base.Server.AuthHead = "HMS 881d3SlFucX5R5hE"
		base.Server.AuthKey = "kHPmWZ4zQBYP24ubmJ5wA4oz0d8EgIFe"
		base.Server.Channel = "100015"
		base.Server.ResUrl = "http://loginios.jr.moefantasy.com/index/getInitConfigs/"
		base.Server.UrlVersion = "http://version.jr.moefantasy.com/index/checkVer/4.1.0/100015/2&version=4.1.0&channel=100015&market=2"
		break
	}
	return base
}
