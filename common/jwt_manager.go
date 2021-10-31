package common

import (
	"fmt"
	"log"
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
	ExtractUserInfo(accessToken string) (jwt.MapClaims, error)
	ExtractTokenMetadata(accessToken string) (jwt.Claims, error)
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
	fmt.Printf("%s", user)
	accessTokenClaims := &JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
			Audience:  user["email"],
		},
		user: user,
	}

	refreshTokenClaims := &JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
		user: user,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	accessTokenString, err := accessToken.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}
	refreshTokenString, err := refreshToken.SignedString([]byte(secret))
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

			return []byte("kushal"), nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("Invalid token.... %w", err)
	}

	_, ok := token.Claims.(*JWTClaims)

	if !ok {
		return nil, fmt.Errorf("Unexpected singing token")
	}

	return map[string]string{
		accessToken: accessToken,
	}, nil
}

// VerifyToken verifies the generated jwt token.
func (jtm *JWTManager) ExtractUserInfo(accessToken string) (jwt.MapClaims, error) {
	hmacSecretString := "kushal"
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	log.Printf("%v", token)

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Printf("Invalid JWT Token:%s\n", ok)
		return nil, fmt.Errorf("%s", "error")
	}

	log.Printf("Invalid JWT Token")
	return claims, nil
}

func (jtm *JWTManager) ExtractTokenMetadata(accessToken string) (jwt.Claims, error) {

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("kushal"), nil
	})
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		log.Printf("Invalid JWT Token:%s\n", ok)

		return nil, err
	}

	if ok && token.Valid {
		if !ok {
			return nil, err
		}
		userId := fmt.Sprintf("%s", claims["email"])
		if err != nil {
			return nil, err
		}
		fmt.Printf("claims:%s", userId)

		return nil, nil
	}

	fmt.Printf("claims:%s", claims)

	return nil, err
}
