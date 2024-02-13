package log_stash

import (
	"github.com/dgrijalva/jwt-go/v4"
)

type JwyPayLoad struct {
	UserName string `json:"userName"`
	NickName string `json:"nickName"`
	RoleID   uint   `json:"roleID"`
	UserID   uint   `json:"userID"`
}

type CustomClaims struct {
	JwyPayLoad
	jwt.StandardClaims
}

// 解析jwt
func ParseToken(token string) (jwyPayLoad *JwyPayLoad) {
	Token, _ := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})
	if Token.Claims == nil {
		return nil
	}
	claims, ok := Token.Claims.(*CustomClaims)
	if !ok {
		// 数据不一致
		return nil
	}
	return &claims.JwyPayLoad
}
