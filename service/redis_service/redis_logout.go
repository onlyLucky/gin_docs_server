package redis_service

import (
	"gin_docs_server/global"
	"time"
)

const prefix = "logout_"

// 设置一个注销token expiration 过期时间
func Logout(token string, expiration time.Duration) error {
	err := global.Redis.Set(prefix+token, "", expiration).Err()
	return err
}

// 判断一个token是否属于注销的token
func CheckLogout(token string) bool {
	_, err := global.Redis.Get(prefix + token).Result()
	if err != nil {
		return false
	}
	return true
}
