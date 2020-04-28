package net

import "fmt"

var commonHeaders = make(map[string]string)

func init() {
	commonHeaders["Accept-Encoding"] = "identity"
	commonHeaders["Connection"] = "Keep-Alive"
	commonHeaders["User-Agent"] = "Dalvik/2.1.0 (Linux; U; Android 5.1.1; mi max Build/LMY48Z)"
}

func FirstLogin(server Server) error {
	fmt.Println(server.UrlVersion)
	data, err := GetGameVersion(server.UrlVersion)
	if err != nil {
		return err
	}
	fmt.Println(data.Version.DataVersion)
	return nil
}
