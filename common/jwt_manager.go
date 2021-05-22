package common

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JWTManager struct {
	secret   string
	duration time.Duration
	user     interface{}
}

type Manager interface {
	VerifyToken(accessToken string) (map[string]string, error)
	GenerateTokens(secret string, duration time.Duration, user interface{}) (map[string]string, error)
}

type JWTClaims struct {
	user interface{}
	jwt.StandardClaims
}

func NewJWTManager() *JWTManager {
	return &JWTManager{}
}

func (jtm *JWTManager) GenerateTokens(secret string, duration time.Duration, user interface{}) (map[string]string, error) {
	accessTokenClaims := &JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jtm.duration).Unix(),
		},
		user: jtm.user,
	}

	refreshTokenClaims := &JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jtm.duration).Unix(),
		},
		user: jtm.user,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	accessTokenString, err := accessToken.SignedString([]byte(jtm.secret))
	if err != nil {
		return nil, err
	}
	refreshTokenString, err := refreshToken.SignedString([]byte(jtm.secret))
	if err != nil {
		return nil, err
	}

	var response = map[string]string{
		"accessToken":  accessTokenString,
		"refreshToken": refreshTokenString,
	}

	return response, nil
}

// VerifyToken verifies the generated jwt token.
func (jtm *JWTManager) VerifyToken(accessToken string) (map[string]string, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("Unexpected singing token")
			}

			return []byte(jtm.secret), nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("Invalid token %w", err)
	}

	_, ok := token.Claims.(*JWTClaims)

	if !ok {
		return nil, fmt.Errorf("Unexpected singing token")
	}

	return map[string]string{
		accessToken: accessToken,
	}, nil
}
