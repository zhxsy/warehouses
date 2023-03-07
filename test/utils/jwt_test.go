package utils

import (
	"github.com/cfx/warehouses/library/utils"
	"github.com/cfx/warehouses/test"
	"github.com/golang-jwt/jwt/v4"
	"testing"
	"time"
)

func TestJwt(t *testing.T) {
	//:= app.Config("app").GetString("signing_key")
	sign := "sign"

	j := &utils.JWT{
		SigningKey: []byte(sign),
	}

	baseclaims := utils.BaseClaims{
		ID:       1,
		NickName: "vic",
		Username: "albert",
	}
	customClaims := j.CreateClaims(baseclaims, 100, jwt.StandardClaims{
		NotBefore: time.Now().Unix() - 1000,   // 签名生效时间
		ExpiresAt: time.Now().Unix() + 604800, // 过期时间 7天  配置文件
		Issuer:    "vic",                      // 签名的发行者
	})
	token, err := j.CreateToken(customClaims)

	test.Printf(t, err, token)

	resp, err := j.ParseToken(token)
	test.Printf(t, err, resp)
}
