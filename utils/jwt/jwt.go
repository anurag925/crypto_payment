package jwt

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtAlgorithm = jwt.SigningMethodHS256
var jwtExpires = 60 * 24 * time.Minute

func Encode(ctx context.Context, claims jwt.MapClaims, secret []byte) (string, error) {
	claims["exp"] = jwt.NewNumericDate(time.Now().Add(jwtExpires))
	claims["iat"] = jwt.NewNumericDate(time.Now())
	claims["iss"] = "crypto_payment"
	claims["sub"] = "authorization"
	// Create a new token object with the claims
	token := jwt.NewWithClaims(jwtAlgorithm, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}
	return tokenString, nil
}

func Decode(ctx context.Context, tokenString string, secret []byte) (jwt.Claims, error) {
	claims := jwt.MapClaims{}
	// Parse the token
	_, err := jwt.ParseWithClaims(
		tokenString, claims, func(token *jwt.Token) (interface{}, error) { return secret, nil })
	return claims, err
}
