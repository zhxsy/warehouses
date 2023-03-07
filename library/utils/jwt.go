package utils

import (
	"errors"
	"github.com/cfx/warehouses/app"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/sync/singleflight"
	"sync"
)

type JWT struct {
	SigningKey []byte    `json:"signing_key"`
	Once       sync.Once `json:"-"`
}

var (
	TokenExpired            = errors.New("Token is expired")
	TokenNotValidYet        = errors.New("Token not active yet")
	TokenMalformed          = errors.New("That's not even a token")
	TokenInvalid            = errors.New("Couldn't handle this token:")
	GVA_Concurrency_Control = &singleflight.Group{}
	OnceJWT                 = &JWT{}
)

func NewJWT() *JWT {
	OnceJWT.Once.Do(func() {
		signingKey := app.Config("app").GetString("signing_key")
		OnceJWT = &JWT{
			SigningKey: []byte(signingKey),
		}
	})
	return OnceJWT
}

// Custom claims structure
type CustomClaims struct {
	BaseClaims         `json:"base_claims"`
	BufferTime         int64 `json:"buffer_time,omitempty"`
	jwt.StandardClaims `json:"jwt_standard_claims"`
}

type BaseClaims struct {
	ID       uint64
	Username string
	NickName string
}

func (j *JWT) CreateClaims(baseClaims BaseClaims, bufferTime int64, standardClaims jwt.StandardClaims) CustomClaims {
	claims := CustomClaims{
		BaseClaims:     baseClaims,
		BufferTime:     bufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: standardClaims,
		//StandardClaims: jwt.StandardClaims{
		//	NotBefore: time.Now().Unix() - 1000,   // 签名生效时间
		//	ExpiresAt: time.Now().Unix() + 604800, // 过期时间 7天  配置文件
		//	Issuer:    "qmPlus",                   // 签名的发行者
		//},
	}
	return claims
}

// 创建一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims CustomClaims) (string, error) {
	v, err, _ := GVA_Concurrency_Control.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}
