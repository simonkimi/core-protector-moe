package net

import "fmt"

/**
1 密码错误
*/
type LoginError struct {
	code    int
	message string
}

func (re *LoginError) Error() string {
	return fmt.Sprintf("登录游戏出现问题 message:%s", re.message)
}
