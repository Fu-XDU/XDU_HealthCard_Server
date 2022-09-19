package auth

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/gommon/random"
	"time"
	"xdu-health-card/utils/flags"
)

type jwtClaims struct {
	OpenId string    `json:"open_id"`
	Iat    time.Time `json:"iat"`
	Exp    time.Time `json:"exp"`
	Rand   string    `json:"rand"`
}

func (j jwtClaims) Valid() error {
	now := time.Now()
	if j.Iat.Before(now) && j.Exp.After(now) {
		return nil
	}
	return errors.New("invalid jwt")
}

func NewJwt(openid string) (tokenStr string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims{
		OpenId: openid,
		Iat:    time.Now(),
		Exp:    time.Now().Add(10 * time.Minute),
		Rand:   random.New().String(31),
	})

	tokenStr, err = token.SignedString([]byte(flags.HmacSecret))
	return
}

func ParseJwt(tokenStr string) (claims *jwtClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(flags.HmacSecret), nil
	})

	if err != nil {
		return
	}
	claims = token.Claims.(*jwtClaims)
	return
}
