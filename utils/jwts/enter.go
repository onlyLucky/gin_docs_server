package jwts

import (
	"github.com/dgrijalva/jwt-go/v4"
)

type JwyPayLoad struct {
	NickName string `json:"nickName"`
	RoleID   uint `json:"roleID"`
	UserID   uint `json:"userID"`
}

var Secret []byte

type CustomClaims struct {
	JwyPayLoad
	jwt.StandardClaims
}