package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/litecodex/go-web-framework/common/utils/crypto/rsa"
	"github.com/litecodex/go-web-framework/web/exceptions"
	"time"
)

type JwtService struct {
	signKey    string // 签名使用
	verifyKey  string // 验签使用
	signMethod string
}

const HS256 string = "HS256"
const RS256 string = "RS256"

func NewJwtService(signMethod, signKey, verifyKey string) *JwtService {
	return &JwtService{
		signMethod: signMethod,
		signKey:    signKey,
		verifyKey:  verifyKey,
	}
}

func (thiz *JwtService) parseSignKey() interface{} {
	signKey := thiz.signKey
	switch thiz.signMethod {
	case HS256:
		return []byte(signKey)
	case RS256:
		return rsa.MustParsePrivateKey(signKey)
	default:
		panic(fmt.Errorf("not supported SignMethod"))
	}
	return nil
}

func (thiz *JwtService) parseVerifyKey() interface{} {
	verifyKey := thiz.verifyKey
	switch thiz.signMethod {
	case HS256:
		return []byte(verifyKey)
	case RS256:
		return rsa.MustParsePublicKey(verifyKey)
	default:
		panic(fmt.Errorf("not supported SignMethod"))
	}
	return nil
}

func (thiz *JwtService) parseSignMethod() jwt.SigningMethod {
	switch thiz.signMethod {
	case HS256:
		return jwt.SigningMethodHS256
	case RS256:
		return jwt.SigningMethodRS256
	default:
		panic(fmt.Errorf("not supported SignMethod"))
	}
	return nil
}

const HOURS = "hours"
const MINUTES = "minutes"
const SECONDS = "seconds"

func (thiz *JwtService) MustCreateToken(payload map[string]interface{}, duration int64, timeUnit string) string {
	token, err := thiz.CreateToken(payload, duration, timeUnit)
	if err != nil {
		panic(err)
	}
	return token
}

// 创建一个JWT并签名
func (thiz *JwtService) CreateToken(payload map[string]interface{}, duration int64, timeUnit string) (string, error) {
	// 设置JWT的声明
	expirationDuration := time.Duration(duration)
	switch timeUnit {
	case HOURS:
		expirationDuration *= time.Hour
	case MINUTES:
		expirationDuration *= time.Minute
	case SECONDS:
		expirationDuration *= time.Second
	default:
		return "", exceptions.OfMessage("invalid timeUnit: must be 'hours', 'minutes', or 'seconds'")
	}

	expirationTime := time.Now().Add(expirationDuration)
	payload["exp"] = expirationTime.Unix()

	claims := jwt.MapClaims(payload)

	// 创建一个token
	token := jwt.NewWithClaims(thiz.parseSignMethod(), claims)
	//token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// 签名并获得完整的编码后的token
	tokenString, err := token.SignedString(thiz.parseSignKey())
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 验证并解析token
func (thiz *JwtService) VerifyAndParseClaims(tokenString string) (map[string]interface{}, error) {
	claims := jwt.MapClaims{}

	// 解析并验证token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return thiz.parseVerifyKey(), nil
	})

	if err != nil {
		// jwt.ErrSignatureInvalid
		// jwt.ErrTokenExpired
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrTokenNotValidYet
	}

	// 返回claims作为map
	payload := make(map[string]interface{})
	for key, value := range claims {
		payload[key] = value
	}

	return payload, nil
}
