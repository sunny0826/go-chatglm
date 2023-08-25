package chatglm

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

const API_TOKEN_TTL_SECONDS = 3 * 60

func generateToken(apikey string) (string, error) {
	parts := strings.Split(apikey, ".")
	if len(parts) != 2 {
		return "", errors.New("invalid apikey")
	}
	id := parts[0]
	secret := parts[1]

	payload := jwt.MapClaims{
		"api_key":   id,
		"exp":       time.Now().Add(time.Second * API_TOKEN_TTL_SECONDS).Unix(),
		"timestamp": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token.Header["alg"] = "HS256"
	token.Header["sign_type"] = "SIGN"

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
