package service

import (
	jwt "github.com/golang-jwt/jwt/v4"
	"monaPanel/global"
	"time"
)

type jwtService struct {
}

var JwtService = new(jwtService)

type JwtUser interface {
	GetUid() string
}

type PanelClaims struct {
	jwt.RegisteredClaims
	// JWT 官方规定的七个字段
	// iss (issuer)：签发人
	// exp (expiration time)：过期时间
	// sub (subject)：主题
	// aud (audience)：受众
	// nbf (Not Before)：生效时间
	// iat (Issued At)：签发时间
	// jti (JWT ID)：编号
	UserId string `json:"user_id"`
}

type claimsToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func (jwtService *jwtService) CreateToken(user JwtUser) (token claimsToken, err error) {
	tokenData := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		PanelClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(global.Config.Jwt.JwtTtl))), // 过期时间
				NotBefore: jwt.NewNumericDate(time.Now()),                                                            // 生效时间
				IssuedAt:  jwt.NewNumericDate(time.Now()),                                                            // 签发时间
			},
			UserId: user.GetUid(),
		},
	)

	tokenStr, err := tokenData.SignedString([]byte(global.Config.Jwt.Secret))

	token = claimsToken{
		AccessToken: tokenStr,
		ExpiresIn:   global.Config.Jwt.JwtTtl,
	}

	return
}
