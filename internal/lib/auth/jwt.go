package auth

import (
	"errors"
	"io/ioutil"
	"time"

	"github.com/go_example/internal/service"

	MysqlModel "github.com/go_example/internal/model/mysql"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type CustomClaims struct {
	UserInfo MysqlModel.User `json:"user_info"`
	jwt.RegisteredClaims
}

const bearerSchema = "Bearer "
const accessTokenExpire = time.Minute * 30
const refreshTokenExpire = time.Hour * 24 * 30

type cert struct {
	PrivateKey []byte
	PublicKey  []byte
}

// 证书密钥
var certs = make(map[string]cert)

// 当前使用的证书key
var curKey string

// SetCerts 生成密钥
func SetCerts() {
	var certKeyConfig struct {
		CurKey string
		Keys   []string
	}
	err := viper.Sub("certs").Unmarshal(&certKeyConfig)
	if err != nil {
		panic(err)
	}
	var privateKey []byte
	var publicKey []byte
	for _, key := range certKeyConfig.Keys {
		privateKey, err = ioutil.ReadFile("assets/certs/" + key + "/rsa_private_key.pem")
		if err != nil {
			panic(err)
		}
		publicKey, err = ioutil.ReadFile("assets/certs/" + key + "/rsa_public_key.pem")
		if err != nil {
			panic(err)
		}
		certs[key] = cert{
			PrivateKey: privateKey,
			PublicKey:  publicKey,
		}
	}

	// 读取当前的使用的证书key
	curKey = viper.GetString("certs.curKey")
}

// GenToken 生成token
func GenToken(authService service.AuthService, userInfo MysqlModel.User) (string, string, error) {
	// 读取当前使用的私钥证书
	secret, err := jwt.ParseRSAPrivateKeyFromPEM(certs[curKey].PrivateKey)
	if err != nil {
		panic(err)
	}
	// 生成access_token
	accessClaims := CustomClaims{
		userInfo,
		jwt.RegisteredClaims{
			Issuer:    "ken",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenExpire)),
		},
	}

	// 使用指定的签名方法创建签名对象
	accessToken := jwt.NewWithClaims(jwt.SigningMethodRS512, accessClaims)

	// 使用指定的secret签名并获得完整的编码后的字符串token
	accessTokenSign, err := accessToken.SignedString(secret)
	if err != nil {
		return "", "", err
	}

	// 生成refresh_token，不包含自定义信息
	refreshClaims := CustomClaims{
		MysqlModel.User{
			Id: userInfo.Id,
		},
		jwt.RegisteredClaims{
			Issuer:    "ken",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshTokenExpire)),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodRS512, refreshClaims)
	refreshTokenSign, err := refreshToken.SignedString(secret)
	if err != nil {
		return "", "", err
	}

	// 将token写入redis hash中
	err = authService.SaveAuthTokenRedis(userInfo.Id, accessTokenSign, refreshTokenSign)
	if err != nil {
		return "", "", err
	}

	return bearerSchema + accessTokenSign, bearerSchema + refreshTokenSign, nil
}

// ParseToken 解析token
func ParseToken(authService service.AuthService, tokenString string, grantType string) (*CustomClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 读取当前使用的公钥证书
		secret, err := jwt.ParseRSAPublicKeyFromPEM(certs[curKey].PublicKey)
		if err != nil {
			panic(err)
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		// 判断token是否有效，若无效返回错误
		err = authService.CheckAuthTokenRedis(claims.UserInfo.Id, tokenString, grantType)
		if err != nil {
			return nil, err
		}
		return claims, nil
	}

	return nil, errors.New("token invalid")
}
