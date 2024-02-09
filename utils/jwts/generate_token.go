package jwts

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"gin_docs_server/global"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

// 生成token
func GenToken(user JwyPayLoad) (string, error) {
	/* secretStr,err:= generateRandomSecret()
	if err != nil {
		logrus.Fatalf("生成jwt 秘钥失败err: %s",err.Error())
	} */
	Secret = []byte(global.Config.Jwt.Secret)
	claims := CustomClaims{
		JwyPayLoad: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Duration(global.Config.Jwt.Expires)*time.Hour)),
			Issuer: global.Config.Jwt.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString(Secret)
}

// 生成一个32位随机数
func generateRandomSecret() (string, error) {
	// 生成32字节的随机数（256位）
	secretBytes := make([]byte, 32)
	if _, err := rand.Read(secretBytes); err != nil {
		return "", fmt.Errorf("failed to read random bytes: %v", err)
	}

	// 将随机字节转换为Base64编码格式，适合用作JWT Secret
	secretBase64 := base64.RawURLEncoding.EncodeToString(secretBytes)

	return secretBase64, nil
}