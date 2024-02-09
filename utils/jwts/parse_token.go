package jwts

import (
	"fmt"

	"github.com/dgrijalva/jwt-go/v4"
)

// 解析jwt
func ParseToken(token string) (*CustomClaims, error) {
	Token,err := jwt.ParseWithClaims(token,&CustomClaims{},func(token *jwt.Token)(interface{},error){
		return []byte("QX0D5UG6SqlW8So324"),nil
	})
	if err != nil {
		return nil,err
	}
	claims,ok := Token.Claims.(*CustomClaims)
	if !ok{
		// 数据不一致
		return nil,err
	}
	if !Token.Valid{
		// 令牌无效
		return nil,err
	}
	fmt.Println("claims,ok:",claims,ok)
	return claims,nil
}